DB_FILE='./sekisan.db'

if [ -e $DB_FILE ]; then
    echo "$DB_FILE already exists. Renaming it."
    DATE=`date "+%Y%m%d_%H%M%S"`
    BACKUP="${DB_FILE}_${DATE}"
    mv $DB_FILE $BACKUP
fi