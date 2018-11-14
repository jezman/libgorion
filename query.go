package libgorion

const (
	queryEvents = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		JOIN Events e ON (e.Event = l.Event)
		JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
		WHERE TimeVal BETWEEN ? AND ?
		AND e.Event BETWEEN 26 AND 29
		ORDER BY TimeVal`
	queryEventsByEmployeeAndDoor = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		JOIN Events e ON (e.Event = l.Event)
		JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
		WHERE TimeVal BETWEEN ? AND ?
		AND e.Event BETWEEN 26 AND 29
		AND p.Name = ?
		AND DoorIndex = ?
		ORDER BY TimeVal`
	queryEventsByEmployee = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		JOIN Events e ON (e.Event = l.Event)
		JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
		WHERE TimeVal BETWEEN ? AND ?
		AND e.Event BETWEEN 26 AND 29
		AND p.Name = ?
		ORDER BY TimeVal`
	queryEventsByDoor = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		JOIN Events e ON (e.Event = l.Event)
		JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
		WHERE TimeVal BETWEEN ? AND ?
		AND e.Event BETWEEN 26 AND 29
		AND DoorIndex = ?
		ORDER BY TimeVal`
	queryEventsDenied = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		JOIN Events e ON (e.Event = l.Event)
		JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
		WHERE TimeVal BETWEEN ? AND ?
		AND e.Event IN (26, 29)
		ORDER BY TimeVal`
	queryWorkedTimeByEmployee = `SELECT p.Name, p.FirstName, p.MidName, c.Name, min(TimeVal), max(TimeVal)
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		WHERE TimeVal BETWEEN ? AND ?
		AND p.Name = ?
		GROUP BY p.Name, p.FirstName, p.MidName, c.Name, CONVERT(varchar(20), TimeVal, 104)`
	queryWorkedTimeByCompany = `SELECT p.Name, p.FirstName, p.MidName, c.Name, min(TimeVal), max(TimeVal)
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		WHERE TimeVal BETWEEN ? AND ?
		AND c.Name = ?
		GROUP BY p.Name, p.FirstName, p.MidName, c.Name, CONVERT(varchar(20), TimeVal, 104)`
	queryWorkedTime = `SELECT p.Name, p.FirstName, p.MidName, c.Name, min(TimeVal), max(TimeVal)
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		WHERE TimeVal BETWEEN ? AND ?
		GROUP BY p.Name, p.FirstName, p.MidName, c.Name, CONVERT(varchar(20), TimeVal, 104)`
	queryEmployees = `SELECT p.Name, p.FirstName, p.MidName, c.Name FROM pList p
		JOIN pCompany c ON (c.ID = p.Company)
		ORDER BY c.Name`
	queryEmployeesByCompany = `SELECT plist.Name, pList.FirstName, pList.MidName, c.Name FROM pList
		JOIN pCompany c ON (c.ID = Company)
		WHERE c.Name = ?
		ORDER BY pList.Name`
	queryCompanies = `SELECT c.Name, Count(pList.Name) FROM pList
		JOIN pCompany c ON (c.ID = Company)
		GROUP BY c.Name`
	queryDoors = "SELECT GIndex, Name FROM AcessPoint ORDER BY GIndex"
	queryEventsValues = "SELECT Event, Contents, Comment from Events"
	queryAddWorker = `DECLARE @pID int
		SELECT @pID = MAX(ID)+1 FROM pList
		INSERT INTO pList (ID, Status, Name, FirstName, MidName)
		VALUES(@pID, 5, ?1, ?2, ?3)`
	queryDeleteWorkerCards = "DELETE FROM pMark WHERE OwnerName = ?1"
	queryDeleteWorker = `DELETE FROM pList
		WHERE Name = ?1
		AND FirstName = ?2
		AND MidName = ?3`
	queryFindWorker = `SELECT ID FROM pList
		WHERE Name = ?1
		AND FirstName = ?2
		AND MidName = ?3`
)
