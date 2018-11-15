package main


type MetaviewEnvSetting struct {
	Logfile       string
	Configfile    string
	Queue         int
	EtcdEndPoints []string
}

envSetting := EnvUtilsParseEnvSetting()
	metaStorageConf := confbase.NewMetaStorageConf(envSetting.EtcdEndPoints)

	/*
	 * Step 2: metaview connect to postgres.
	 * if db connection failed, would keep retry, always at step 2.
	 */
	err, metaEtcdConfig := metaStorageConf.LoadMetaEtcdConfig(envSetting.Queue)
	if err != nil {
		Logger.Errorf("Can't load configuration from etcd: %s\n",
			err.Error())
		return err
	}

	queueDBConfig := &metaEtcdConfig.DBConfig
	dbService := dbservice.NewDBService(queueDBConfig)
	if dbService == nil {
		Logger.Errorf("Fail to init DB service")
		return errors.New("Fail to init database")
	}



func EnvUtilsParseEnvSetting() *MetaviewEnvSetting {
	envSetting := &MetaviewEnvSetting{
		Logfile:    "/var/log/metaview/metaview.log",
		Configfile: "metaview.json",
	}
	logFile := os.Getenv("LOGFILE")
	if logFile != "" {
		envSetting.Logfile = logFile
	}
	configFile := os.Getenv("CONFIGFILE")
	if configFile != "" {
		envSetting.Configfile = configFile
	}
	etcdHosts := os.Getenv("ETCD_CLUSTER")
	if etcdHosts == "" {
		etcdHosts = "http://localhost:2379"
	}
	queue := os.Getenv("QUEUE")
	if queue == "" {
		envSetting.Queue = 0
	} else {
		queueNum, err := strconv.Atoi(queue)
		if err != nil {
			Logger.Errorf("Queue env %s set wrong: %s\n",
				queue, err.Error())
			envSetting.Queue = 0
		} else {
			envSetting.Queue = queueNum
		}
	}
	hostItems := strings.Split(etcdHosts, ",")
	endPoints := make([]string, 0, len(hostItems))
	for i := 0; i < len(hostItems); i++ {
		if hostItems[i] != "" {
			endPoints = append(endPoints, hostItems[i])
		}
	}

	envSetting.EtcdEndPoints = endPoints
	return envSetting
}
