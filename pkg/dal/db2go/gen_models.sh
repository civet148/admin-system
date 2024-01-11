@echo off


OUT_DIR=..
PACK_NAME=models
SUFFIX_NAME="do"
READ_ONLY="created_time,updated_time"
DB_NAME="admin-system"
WITH_OUT=""
TAGS="bson"
DSN_URL="mysql://root:123456@192.168.1.16:3306/admin-system?charset=utf8"
JSON_PROPERTIES="omitempty"
SPEC_TYPES=""
TINYINT_TO_BOOL="deleted,disabled,ok,is_admin,is_inherent,is_offline,is_default"
TABLE_NAME=""
IMPORT_MODELS=admin-system/pkg/dal/models

db2go --url "${DSN_URL}" --out "${OUT_DIR}" --db "${DB_NAME}" --table "${TABLE_NAME}" --enable-decimal --spec-type "${SPEC_TYPES}" \
      --suffix "${SUFFIX_NAME}" --package "${PACK_NAME}" --readonly "${READ_ONLY}" --without "${WITH_OUT}" --tag "${TAGS}" --tinyint-as-bool "${TINYINT_TO_BOOL}" \
      --dao dao --import-models "${IMPORT_MODELS}"

echo generate go file ok, formatting...
gofmt -w %OUT_DIR%/%PACK_NAME%

