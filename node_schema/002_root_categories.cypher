MATCH (root:Node {name:"ROOT"})

CREATE (olymp:Node {name:"Олимпиады и конкурсы", points:0})
CREATE (clubs:Node {name:"Кружки", points:0})
CREATE (shifts:Node {name:"Профильные смены", points:0})
CREATE (awards:Node {name:"Награды", points:0})

CREATE (root)-[:NEXT]->(olymp)
CREATE (root)-[:NEXT]->(clubs)
CREATE (root)-[:NEXT]->(shifts)
CREATE (root)-[:NEXT]->(awards);
