## Overview

This is an utils package for multiple usages. 
It contains basic utility functions for strings, mysql, mongo, elastic-search, files etc.


### Library

    1.  MYSQL Package
```go
package main

import (
	"github.com/zapr-oss/go-utils/mysql"
	"github.com/zapr-oss/go-utils/mysql/config"
)

func main() {
	mConfig := mysqlconfig.Config{
		Host: "127.0.0.1",
		User: "username",
		Password:"password",
		Database:"database",
		Port:3306,
		RetryCount: 4,
	}


	mysqlE, err := mysql.Connect(mConfig)
	if err != nil {
		// handle error
	}
	defer mysqlE.Close()

	rows, err := mysqlE.Query("Select * from apples")

	if err != nil {
	//	handle error
	}
    
	for rows.Next() {
	//	iter rows
	}
    //iterate through rows.
}
```

    2. MONGO Package
    

```go
package main

import (
	"github.com/zapr-oss/go-utils/mongo"
	"github.com/zapr-oss/go-utils/mongo/config"
	"log"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	config := mongoconfig.Config{
        Host: "127.0.0.1",
        UserName: "username",
        Password:"password",
        Database:"database",
        Port:3306,
    }

	mongoE, err := mongo.Connect(config)
	if err != nil {
	    //Handle error
	}
	defer mongoE.Close()

	col := mongoE.DB.C("collection_name")

	columnCount, _ := col.Find(bson.M{}).Count()

	log.Println(columnCount)
}

```

    3. GRAPHITE Package
```go
package main

import (
    "github.com/zapr-oss/go-utils/graphite/graphiteutil"
    "github.com/zapr-oss/go-utils/graphite"
)

func main() {

    graphiteConfig := graphite.GraphiteConfig{
       Host: "127.0.0.1",
       Port:3306,
       Prefix: "metric_prefix",
       Environment: "prod",
       FlushIntervalInSec: 1, // interval to send metrics to graphite server.
       Disabled: false,
   }

    err := graphiteutil.InitializeGraphite(graphiteConfig)

    if err != nil {
    //	Handle error
    }

    graphite.GetCounter("counter_name").Inc()
}

```

    4. ElasticSearch Package
```go
package main

import (
	"github.com/zapr-oss/go-utils/elasticsearch/config"
    "github.com/zapr-oss/go-utils/elasticsearch"
)
func main() {
	elasticConfig := elasticconfig.Config{Addresses: []string{"http://127.0.0.1:9200", "http://192.168.0.1:9200"}}
	client, err := elasticsearch.Connect(elasticConfig)
	
	if err != nil {
	// Handle error
	}
}

```

Except the above functionality there are some utils functions for multiple purposes as mentioned above.