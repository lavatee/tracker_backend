MATCH (olymp:Node {name:"Олимпиады и конкурсы"})
MATCH (clubs:Node {name:"Кружки"})
MATCH (shifts:Node {name:"Профильные смены"})
MATCH (awards:Node {name:"Награды"})

CREATE (c1:Node {name:"Критерий 1: Участие в профильных мероприятиях", points:0})
CREATE (c2:Node {name:"Критерий 2: Участие в НТО", points:0})
CREATE (c3:Node {name:"Критерий 3: Участие во Всероссийской олимпиаде школьников", points:0})
CREATE (c4:Node {name:"Критерий 4: АгроНТРИ и Большие вызовы", points:0})
CREATE (c6:Node {name:"Критерий 6: Научно-практические конференции", points:0})

CREATE (olymp)-[:NEXT]->(c1)
CREATE (olymp)-[:NEXT]->(c2)
CREATE (olymp)-[:NEXT]->(c3)
CREATE (olymp)-[:NEXT]->(c4)
CREATE (olymp)-[:NEXT]->(c6)

CREATE (c7:Node {name:"Критерий 7: Посещение профильных кружков", points:0})
CREATE (clubs)-[:NEXT]->(c7)

CREATE (c5:Node {name:"Критерий 5: Профильные смены (Альтаир/Сириус)", points:0})
CREATE (shifts)-[:NEXT]->(c5)

CREATE (c8:Node {name:"Критерий 8: Приглашение нового пользователя", points:0})
CREATE (c9:Node {name:"Критерий 9: Отчетность о развитии проекта", points:0})
CREATE (c10:Node {name:"Критерий 10: Своевременное прикрепление наградных материалов", points:0})
CREATE (c11:Node {name:"Критерий 11: За содействие активу", points:0})

CREATE (awards)-[:NEXT]->(c8)
CREATE (awards)-[:NEXT]->(c9)
CREATE (awards)-[:NEXT]->(c10)
CREATE (awards)-[:NEXT]->(c11);
