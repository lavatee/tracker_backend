package service

type UsersService struct {
	repo *repository.Repository
}

func NewUsersService(repo *repository.Repository) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) hashPassword(password string) string {
	sha := sha1.New()
	sha.Write([]byte(password))
	return fmt.Sprintf("%x", sha.Sum([]byte(salt)))
}

func (s *UsersService) SignUp(user model.User) (int, error) {
	user.PasswordHash = s.hashPassword(user.PasswordHash)
	referral, err := uuid.GenerateUUID()
	if err != nil {
		return 0, err
	}
	user.Referral = referral
	return s.repo.Users.SignUp(user)
}

func (s *UsersService) NewToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString([]byte(tokenKey))
	if err != nil {
		return "", err
	}
	return stringToken, nil
}

func (s *UsersService) SignIn(telegramUsername string, password string) (string, string, error) {
	user, err := s.repo.Users.SignIn(telegramUsername, s.hashPassword(password))
	if err != nil {
		return "", "", err
	}
	userId := user.Id
	accessClaims := jwt.MapClaims{
		"exp": time.Now().Add(accessTTL).Unix(),
		"id":  userId,
	}
	refreshClaims := jwt.MapClaims{
		"exp": time.Now().Add(refreshTTL).Unix(),
		"id":  userId,
	}
	access, err := s.NewToken(accessClaims)
	if err != nil {
		return "", "", err
	}
	refresh, err := s.NewToken(refreshClaims)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

func (s *UsersService) Refresh(refreshToken string) (string, string, error) {
	parsedToken, err := jwt.ParseWithClaims(refreshToken, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token is invalid")
		}
		return []byte(tokenKey), nil
	})
	if err != nil {
		return "", "", err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		accessClaims := jwt.MapClaims{
			"exp": time.Now().Add(accessTTL).Unix(),
			"id":  claims["id"],
		}
		refreshClaims := jwt.MapClaims{
			"exp": time.Now().Add(refreshTTL).Unix(),
			"id":  claims["id"],
		}
		access, err := s.NewToken(accessClaims)
		if err != nil {
			return "", "", err
		}
		refresh, err := s.NewToken(refreshClaims)
		if err != nil {
			return "", "", err
		}
		return access, refresh, nil
	}
	return "", "", fmt.Errorf("token is invalid")
}

func (s *UsersService) ParseToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token is invalid")
		}
		return []byte(tokenKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("token is expired")
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token is expired")
}

func (s *UsersService) GetOneUser(userId int) (model.User, error) {
	return s.repo.Users.GetOneUser(userId)
}

