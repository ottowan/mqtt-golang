# mqtt-golang

    #implement by
        #golang
        #mqtt
        #broker : hivemq
    
#Run Subscribe
    cd mqtt_sub
    go run main.go osub01 "ottowan/light"

    #comment 
    go run main.go <subscribe_name> <topic>
    
    
#Run Publisher 
    cd mqtt_pub
    go run main.go opub01 "ottowan/light" "Hello light"

    #comment 
    go run main.go <publisher_name> <topic> <message>
