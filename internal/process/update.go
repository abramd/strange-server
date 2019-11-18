package process

import (
	"fmt"
	"github.com/abramd/strange-server/internal/model"
	"log"
	"time"
)

var (
	updateStateDeletedQuery = fmt.Sprintf(`UPDATE %s SET state='%s' WHERE is_processed=true`, model.EntityTable, model.StateDeleted)
	updateProcessedQuery    = fmt.Sprintf(`WITH a AS (SELECT id, is_processed FROM %s WHERE is_processed=false ORDER BY id ASC LIMIT 10) UPDATE %s SET is_processed=true FROM a`, model.EntityTable, model.EntityTable)
)

func UpdateStateDeleted(m *model.Manager, dur time.Duration) {
	t := time.NewTicker(dur)
	for {
		select {
		case <-t.C:
			log.Println("update deleted: error:", updateStateDeleted(m))
		}
	}
}

func UpdateProcessed(m *model.Manager, dur time.Duration) {
	t := time.NewTicker(dur)
	for {
		select {
		case <-t.C:
			log.Println("update processed: error:", updateProcessed(m))
		}
	}
}

func updateStateDeleted(m *model.Manager) error {
	return m.ExecTx(updateStateDeletedQuery)
}

func updateProcessed(m *model.Manager) error {
	return m.ExecTx(updateProcessedQuery)
}
