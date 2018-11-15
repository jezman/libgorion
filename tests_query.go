package libgorion

const (
	testQueryEmployees = `^SELECT (.+), (.+), (.+), (.+) FROM pList p
		JOIN pCompany c ON ((.+) = (.+))
		ORDER BY (.+)$`
	testQueryEmployeesByCompany = `^SELECT (.+), (.+), (.+), (.+) FROM pList
		JOIN pCompany c ON ((.+) = (.+))
		WHERE c.Name = (.+)
		ORDER BY (.+)$`
	testQueryCompanies = `^SELECT (.+), (.+) FROM pList
		JOIN pCompany c ON ((.+) = (.+))
		GROUP BY (.+)$`
	testQueryDoors  = "^SELECT (.+), (.+) FROM AcessPoint ORDER BY GIndex$"
	testQueryEvents = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+)
	    FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		JOIN Events e ON ((.+) = (.+))
		JOIN AcessPoint a ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		AND (.+) BETWEEN 26 AND 29
		ORDER BY (.+)$`
	testQueryEventsByEmployeeAndDoor = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		JOIN Events e ON ((.+) = (.+))
		JOIN AcessPoint a ON ((.+) = (.+))
		WHERE (.+) BETWEEN (.+) AND (.+)
		AND e.Event BETWEEN 26 AND 29
		AND p.Name = (.+)
		AND DoorIndex = (.+)
		ORDER BY (.+)$`
	testQueryEventsByEmployee = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		JOIN Events e ON ((.+) = (.+))
		JOIN AcessPoint a ON ((.+) = (.+))
		WHERE (.+) BETWEEN (.+) AND (.+)
		AND e.Event BETWEEN 26 AND 29
		AND p.Name = (.+)
		ORDER BY (.+)$`
	testQueryEventsByDoor = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		JOIN Events e ON ((.+) = (.+))
		JOIN AcessPoint a ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		AND e.Event BETWEEN 26 AND 29
		AND DoorIndex = (.+)
		ORDER BY (.+)$`
	testQueryEventsDenied = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+)
	    FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		JOIN Events e ON ((.+) = (.+))
		JOIN AcessPoint a ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		AND (.+) IN (.+)
		ORDER BY (.+)$`
	testQueryWorkedTime = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		GROUP BY (.+), (.+), (.+), (.+), (.+)$`
	testQueryWorkedTimeByCompany = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		AND (.+) = (.+)
		GROUP BY (.+), (.+), (.+), (.+), (.+)$`
	testQueryWorkedTimeByEmployee = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		AND (.+) = (.+)
		GROUP BY (.+), (.+), (.+), (.+), (.+)$`
	testQueryEventsValues       = "SELECT (.+), (.+), (.+) from Events"
	testQueryFindWorkerIDByName = `^SELECT (.+) FROM pList
		WHERE Name = (.+)
		AND FirstName = (.+)
		AND MidName = (.+)$`
	testQueryAddWorker = `^DECLARE (.+) int
		SELECT (.+) = (.+) FROM pList
		INSERT INTO pList (ID, Status, Name, FirstName, MidName)
		VALUES((.+), 5, (.+), (.+), (.+))`
	testQueryDeleteWorkerCards = "^DELETE FROM pMark WHERE OwnerName = (.+)$"
	testQueryDeleteWorker      = `^DELETE FROM pList
		WHERE Name = (.+)
		AND FirstName = (.+)
		AND MidName = (.+)$`
	testQueryUpdateWorkerCardStatus = `^UPDATE pMark SET Config = (.+) WHERE OwnerName = (.+)$`
)
