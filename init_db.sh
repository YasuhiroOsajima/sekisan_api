DB_FILE='./sekisan.sqlite3'

# Backup old database file.

if [ -e ${DB_FILE} ]; then
    echo "${DB_FILE} already exists. Renaming it."
    DATE=`date "+%Y%m%d_%H%M%S"`
    BACKUP="${DB_FILE}_${DATE}"
    mv ${DB_FILE} ${BACKUP}
fi


# Create new database file.

echo ".open ${DB_FILE}" | sqlite3

if [ ! -e ${DB_FILE} ]; then
    echo "Creating ${DB_FILE} is failed."
    exit 2
fi

SQLITE_COM="sqlite3 ${DB_FILE} "


# Create tables.

echo "Create 'admin' table."

${SQLITE_COM} "CREATE TABLE IF NOT EXISTS admin (\
  name        TEXT  PRIMARY KEY,\
  password    TEXT,\
  enabled     INT\
  );"


echo "Create 'member' table."

${SQLITE_COM} "CREATE TABLE IF NOT EXISTS member (\
  employee_num  INT  PRIMARY KEY,\
  name          TEXT,\
  enabled       INT\
  );"


echo "Create 'sekisan' table."

${SQLITE_COM} "CREATE TABLE IF NOT EXISTS sekisan (\
  id            INTEGER PRIMARY KEY  AUTOINCREMENT,\
  employee_num  INT,\
  sekisan       INT,\
  FOREIGN KEY (employee_num)\
  REFERENCES member(employee_num)\
  );"


echo "Create index on 'sekisan' table."

${SQLITE_COM} "CREATE INDEX IF NOT EXISTS sekisan_idx \
  ON sekisan(employee_num);"


echo "Create 'transactions' table."

${SQLITE_COM} "CREATE TABLE IF NOT EXISTS transactions (\
  id            INTEGER PRIMARY KEY  AUTOINCREMENT,\
  sekisan_id    INT,\
  update_date   TEXT,\
  employee_num  INT,\
  before        INT,\
  added         INT,\
  subtracted    INT,\
  after         INT,\
  reason        TEXT,\
  FOREIGN KEY (sekisan_id)\
  REFERENCES sekisan(id),\
  FOREIGN KEY (employee_num)\
  REFERENCES member(employee_num)\
  );"


echo "Create index on 'transactions' table."

${SQLITE_COM} "CREATE INDEX IF NOT EXISTS transaction_id_emp_idx \
  ON transactions(id, employee_num);"


# Insert records for debugging.

echo "Insert debug records to 'sekisan' table."

${SQLITE_COM} "INSERT INTO sekisan(employee_num, sekisan) VALUES (2001, 10);"
	//h := controller.NewHandler(db, store)
${SQLITE_COM} "INSERT INTO sekisan(employee_num, sekisan) VALUES (2002, 11);"
