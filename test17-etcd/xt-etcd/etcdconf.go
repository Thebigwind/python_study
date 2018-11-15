package confbase

import (
	"encoding/json"
	"fmt"

	. "github.com/xtao/metaview/common"
	xconfig "github.com/xtao/xstone/config"
)

const (
	ETCD_KEY_CONFIG_DB = "/datamgmt/metaview"
)

type MetaEtcdConfig struct {
	DBConfig   DBServiceConfig
	RestConfig RestServerConfig
}

type MetaStorageConf struct {
	xconfig.StorageMgrConfDB
}

func NewMetaStorageConf(endpoints []string) *MetaStorageConf {
	client := &MetaStorageConf{
		StorageMgrConfDB: *xconfig.NewStorageMgrConfDB(endpoints),
	}

	return client
}

func (db *MetaStorageConf) LoadMetaEtcdConfig(queue int) (error, MetaEtcdConfig) {
	queueKey := fmt.Sprintf("%d", queue)

	etcdConf := MetaEtcdConfig{}
	err, rawData := db.GetSingleKey(ETCD_KEY_CONFIG_DB + "/" + queueKey)
	if err != nil {
		Logger.Errorf("Can't get DBConfig %s",
			ETCD_KEY_CONFIG_DB)
		return err, etcdConf
	}

	err = json.Unmarshal(rawData, &etcdConf)
	if err != nil {
		Logger.Errorf("Can't decode DB etcd config: %s\n",
			err.Error())
		return err, etcdConf
	}

	Logger.Infof("get etcdConf: %s\n", etcdConf)
	return nil, etcdConf
}
