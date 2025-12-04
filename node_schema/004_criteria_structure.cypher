MATCH (c1:Node {name:"Критерий 1: Участие в профильных мероприятиях"})

CREATE (level_regional:Node {name:"Отборочный этап (регион)", points:4})
CREATE (level_municipal:Node {name:"Отборочный этап (муниципальный)", points:3})

CREATE (final_all_rus:Node {name:"Финальные этапы (всероссийский)", points:5})
CREATE (final_regional:Node {name:"Финальные этапы (регион)", points:4})
CREATE (final_municipal:Node {name:"Финальные этапы (муниципальный)", points:3})

CREATE (winner_all_rus:Node {name:"Победитель/призер (всеросс., межрег.)", points:15})
CREATE (winner_all_rus_low:Node {name:"Призер (всеросс., межрег.)", points:12})
CREATE (winner_regional:Node {name:"Победитель/призер (регион)", points:10})
CREATE (winner_regional_low:Node {name:"Призер (регион)", points:7})
CREATE (winner_municipal:Node {name:"Победитель/призер (муницип)", points:5})
CREATE (winner_municipal_low:Node {name:"Призер (муницип)", points:3})

CREATE (c1)-[:NEXT]->(level_regional)
CREATE (c1)-[:NEXT]->(level_municipal)

CREATE (c1)-[:NEXT]->(final_all_rus)
CREATE (c1)-[:NEXT]->(final_regional)
CREATE (c1)-[:NEXT]->(final_municipal)

CREATE (c1)-[:NEXT]->(winner_all_rus)
CREATE (c1)-[:NEXT]->(winner_all_rus_low)
CREATE (c1)-[:NEXT]->(winner_regional)
CREATE (c1)-[:NEXT]->(winner_regional_low)
CREATE (c1)-[:NEXT]->(winner_municipal)
CREATE (c1)-[:NEXT]->(winner_municipal_low);



MATCH (c2:Node {name:"Критерий 2: Участие в НТО"})

CREATE (junior:Node {name:"НТО Junior (7 класс)", points:0})
CREATE (nto8:Node {name:"НТО 8–11 классы", points:0})

CREATE (c2)-[:NEXT]->(junior)
CREATE (c2)-[:NEXT]->(nto8)

CREATE (jun1:Node {name:"Отборочный этап (Junior)", points:3})
CREATE (jun2:Node {name:"Заключительный этап (Junior)", points:5})
CREATE (jun_win:Node {name:"Победитель/призер (Junior)", points:10})

CREATE (junior)-[:NEXT]->(jun1)
CREATE (junior)-[:NEXT]->(jun2)
CREATE (junior)-[:NEXT]->(jun_win)

CREATE (step1:Node {name:"Первый отборочный этап", points:0})
CREATE (step2:Node {name:"Второй отборочный этап", points:10})
CREATE (final:Node {name:"Заключительный этап", points:20})
CREATE (win:Node {name:"Победа", points:30})

CREATE (nto8)-[:NEXT]->(step1)
CREATE (nto8)-[:NEXT]->(step2)
CREATE (nto8)-[:NEXT]->(final)
CREATE (nto8)-[:NEXT]->(win)

CREATE (step1_subject:Node {name:"Предметный тур", points:3})
CREATE (step1_engineer:Node {name:"Инженерный тур", points:7})

CREATE (step1)-[:NEXT]->(step1_subject)
CREATE (step1)-[:NEXT]->(step1_engineer);



MATCH (c3:Node {name:"Критерий 3: Участие во Всероссийской олимпиаде школьников"})

CREATE (school:Node {name:"Школьный этап", points:0})
CREATE (mun:Node {name:"Муниципальный этап", points:0})
CREATE (reg:Node {name:"Региональный этап", points:0})
CREATE (allrus:Node {name:"Всероссийский этап", points:0})

CREATE (c3)-[:NEXT]->(school)
CREATE (c3)-[:NEXT]->(mun)
CREATE (c3)-[:NEXT]->(reg)
CREATE (c3)-[:NEXT]->(allrus)

CREATE (school_win:Node {name:"Победитель школьного этапа", points:3})
CREATE (school_prize:Node {name:"Призер школьного этапа", points:1})

CREATE (mun_win:Node {name:"Победитель муниципального этапа", points:7})
CREATE (mun_prize:Node {name:"Призер муниципального этапа", points:4})

CREATE (reg_win:Node {name:"Победитель регионального этапа", points:15})
CREATE (reg_prize:Node {name:"Призер регионального этапа", points:10})

CREATE (allrus_win:Node {name:"Победитель всероссийского этапа", points:25})
CREATE (allrus_prize:Node {name:"Призер всероссийского этапа", points:20})

CREATE (school)-[:NEXT]->(school_win)
CREATE (school)-[:NEXT]->(school_prize)
CREATE (mun)-[:NEXT]->(mun_win)
CREATE (mun)-[:NEXT]->(mun_prize)
CREATE (reg)-[:NEXT]->(reg_win)
CREATE (reg)-[:NEXT]->(reg_prize)
CREATE (allrus)-[:NEXT]->(allrus_win)
CREATE (allrus)-[:NEXT]->(allrus_prize);



MATCH (c4:Node {name:"Критерий 4: АгроНТРИ и Большие вызовы"})

CREATE (extr1:Node {name:"Заочный региональный этап", points:15})
CREATE (extr2:Node {name:"Очный региональный этап", points:20})
CREATE (extr3:Node {name:"Заключительный этап", points:25})

CREATE (c4)-[:NEXT]->(extr1)
CREATE (c4)-[:NEXT]->(extr2)
CREATE (c4)-[:NEXT]->(extr3);



MATCH (c5:Node {name:"Критерий 5: Профильные смены (Альтаир/Сириус)"})

CREATE (alt:Node {name:"Альтаир", points:10})
CREATE (sirius:Node {name:"Сириус", points:20})

CREATE (c5)-[:NEXT]->(alt)
CREATE (c5)-[:NEXT]->(sirius);



MATCH (c6:Node {name:"Критерий 6: Научно-практические конференции"})

CREATE (okr:Node {name:"Окружной уровень", points:0})
CREATE (mun:Node {name:"Муниципальный уровень", points:0})
CREATE (reg:Node {name:"Региональный уровень", points:0})
CREATE (allrus:Node {name:"Всероссийский уровень", points:0})

CREATE (c6)-[:NEXT]->(okr)
CREATE (c6)-[:NEXT]->(mun)
CREATE (c6)-[:NEXT]->(reg)
CREATE (c6)-[:NEXT]->(allrus)

CREATE (okr_win:Node {name:"Окружной победитель", points:10})
CREATE (okr_prize:Node {name:"Окружной призер", points:7})
CREATE (okr_part:Node {name:"Окружной участник", points:3})
CREATE (okr)-[:NEXT]->(okr_win)
CREATE (okr)-[:NEXT]->(okr_prize)
CREATE (okr)-[:NEXT]->(okr_part)

CREATE (mun_win:Node {name:"Муниципальный победитель", points:12})
CREATE (mun_prize:Node {name:"Муниципальный призер", points:10})
CREATE (mun_part:Node {name:"Муниципальный участник", points:7})
CREATE (mun)-[:NEXT]->(mun_win)
CREATE (mun)-[:NEXT]->(mun_prize)
CREATE (mun)-[:NEXT]->(mun_part)

CREATE (reg_win:Node {name:"Региональный победитель", points:15})
CREATE (reg_prize:Node {name:"Региональный призер", points:12})
CREATE (reg_part:Node {name:"Региональный участник", points:10})
CREATE (reg)-[:NEXT]->(reg_win)
CREATE (reg)-[:NEXT]->(reg_prize)
CREATE (reg)-[:NEXT]->(reg_part)

CREATE (all_win:Node {name:"Всероссийский победитель", points:20})
CREATE (all_prize:Node {name:"Всероссийский призер", points:15})
CREATE (all_part:Node {name:"Всероссийский участник", points:10})
CREATE (allrus)-[:NEXT]->(all_win)
CREATE (allrus)-[:NEXT]->(all_prize)
CREATE (allrus)-[:NEXT]->(all_part);



MATCH (c7:Node {name:"Критерий 7: Посещение профильных кружков"})

CREATE (lyc:Node {name:"Лицей", points:3})
CREATE (kvant:Node {name:"Кванториум / IT-центр", points:5})
CREATE (univ:Node {name:"Университетские кружки (НГТУ, НГПУ, Яндекс, Тинькофф)", points:7})
CREATE (online_reg:Node {name:"Онлайн-курс (региональный оператор)", points:5})
CREATE (online_fed:Node {name:"Онлайн-курс (федеральный оператор)", points:5})

CREATE (c7)-[:NEXT]->(lyc)
CREATE (c7)-[:NEXT]->(kvant)
CREATE (c7)-[:NEXT]->(univ)
CREATE (c7)-[:NEXT]->(online_reg)
CREATE (c7)-[:NEXT]->(online_fed);



MATCH (c8:Node {name:"Критерий 8: Приглашение нового пользователя"})

CREATE (inviter:Node {name:"Пригласитель", points:3})
CREATE (invited:Node {name:"Приглашенный", points:2})

CREATE (c8)-[:NEXT]->(inviter)
CREATE (c8)-[:NEXT]->(invited);



MATCH (c9:Node {name:"Критерий 9: Отчетность о развитии проекта"})

CREATE (rep:Node {name:"Отчетность за месяц", points:1})
CREATE (c9)-[:NEXT]->(rep);



MATCH (c10:Node {name:"Критерий 10: Своевременное прикрепление наградных материалов"})

CREATE (mat:Node {name:"Прикрепление наградных материалов", points:1})
CREATE (c10)-[:NEXT]->(mat);



MATCH (c11:Node {name:"Критерий 11: За содействие активу"})

CREATE (help:Node {name:"Помощь в проведении мероприятий", points:1})
CREATE (c11)-[:NEXT]->(help);
