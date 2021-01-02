
# capture-the-flag-mongodb-go-server

	Backend code that's written in Golang and built with MongoDB and GraphQL.

Build Cmd

    go build capture-the-flag.com/m

Dependencies

	"github.com/graph-gophers/graphql-go"
    "go.mongodb.org/mongo-driver"
    "github.com/zebresel-com/mongodm"
    "gopkg.in/mgo.v2"
    

API (Query and Mutation) Methods

    # Gets up to 1000000 players (with offset = 0)
    {
        players(first: 1000000, offset: 0) {
            id
                firstname
                lastname
                location
                    {
                    latitude
                    longitude
                    }
            }
    }

    # Gets a specific player
    {
        player(id: "5feebfdfc54d2d104fdf6ff9") {
            firstname
            lastname
            location {
                longitude
                latitude
            }
        }
    }

    # Create a player (returns id, firstname, location if successful)
    mutation createPlayer {
        createPlayer(input: 
            {player: 
                {
                    firstname : "player_first_name_string", 
                    lastname  : "player_last_name_string",

                     location : 
                        {
                        longitude : 123.23123123,
                        latitude  : 234.123123123
                        }
                }
            }
        ) 
        {
            id
            firstname
          	lastname
            location {
            longitude
            latitude
            }
        }
    }



    # Updates a player (returns id, firstname, location if successful)
    # This mutation is currently under construction
    mutation updatePlayer {
        updatePlayer(
            input: {
                id: "5feeb0a8c54d2d0e900984a9", 
                patch: {
                    firstname : "new_first_name_string", 
                    lastname  : "new_last_name_string",

                       location : 
                          {
                              longitude : 199.23123123,
                              latitude  : 234.123123123
                          }
                      }
            })
        {
            id
            firstname
            lastname
            location {
            longitude
            latitude
            }
        }
    }



    # Deletes a player
    mutation deletePlayer {
        deletePlayer(input: {id: "5ff09890c54d2d1e2be59319"}) {
            firstname
        }
    }

Instructions 

    