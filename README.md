## Overview

This is an utils package for multiple usages. 
It contains basic utility functions for strings, mysql, mongo, elastic-search, files etc.


### Library

    1.  MYSQL Package
```go
package main

import (
	"bitbucket.org/zapr/go-utils/mysql"
	"bitbucket.org/zapr/go-utils/mysql/config"
)

func main() {
	mConfig := mysqlconfig.Config{
		Host: "102.0.0.1",
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
	"bitbucket.org/zapr/go-utils/mongo"
	"bitbucket.org/zapr/go-utils/mongo/config"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	config := mongoconfig.Config{
	    //provide config
	}

	mongoE, err := mongo.Connect(config)
	if err != nil {
	    //Handle error
	}
	defer mongoE.Close()

	col := mongoE.DB.C("collection_name")

	columnCount, _ := col.Find(bson.M{}).Count()

	fmt.Println(columnCount)
}

```

    3. GRAPHITE Package
```go
    package main
    
    import (
    	"bitbucket.org/zapr/go-utils/graphite/graphiteutil"
    	"bitbucket.org/zapr/go-utils/graphite"
    )
    
    func main() {
    
    	graphiteConfig := graphite.GraphiteConfig{}
    
    	err := graphiteutil.InitializeGraphite(graphiteConfig)
    
    	if err != nil {
    	//	Handle error
    	}
    
    	graphite.GetCounter("counter_name").Inc()
    }

```

Except the above functionality there are some utils functions for multiple purposes as mentioned above.