module github.com/cgrates/cgrates

go 1.18

// replace github.com/cgrates/radigo => ../radigo

// replace github.com/cgrates/rpcclient => ../rpcclient

require (
	github.com/Azure/go-amqp v0.15.0
	github.com/antchfx/xmlquery v1.3.3
	github.com/aws/aws-sdk-go v1.36.24
	github.com/blevesearch/bleve v1.0.14
	github.com/cenkalti/rpc2 v0.0.0-20210220005819-4a29bc83afe1
	github.com/cgrates/aringo v0.0.0-20201113143849-3b299e4e636d
	github.com/cgrates/baningo v0.0.0-20201105145354-6e3173f6a91b
	github.com/cgrates/fsock v0.0.0-20191107070144-e7a331109df7
	github.com/cgrates/kamevapi v0.0.0-20191001125829-7dbc3ad58817
	github.com/cgrates/ltcache v0.0.0-20181016092649-92fb7fa77cca
	github.com/cgrates/radigo v0.0.0-20201113143731-162035428d72
	github.com/cgrates/rpcclient v0.0.0-20210218104959-cc39fa26221e
	github.com/cgrates/sipingo v1.0.1-0.20200514112313-699ebc1cdb8e
	github.com/cgrates/ugocodec v0.0.0-20201023092048-df93d0123f60
	github.com/creack/pty v1.1.11
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/elastic/go-elasticsearch v0.0.0
	github.com/ericlagergren/decimal v0.0.0-20191206042408-88212e6cfca9
	github.com/fiorix/go-diameter/v4 v4.0.2
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	github.com/mediocregopher/radix/v3 v3.7.0
	github.com/miekg/dns v1.1.44-0.20210927135021-1630ffe2ca11
	github.com/mitchellh/mapstructure v1.4.0
	github.com/nats-io/nats.go v1.11.0
	github.com/nyaruka/phonenumbers v1.0.73
	github.com/peterh/liner v1.2.1
	github.com/segmentio/kafka-go v0.4.8
	github.com/streadway/amqp v1.0.0
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/crypto v0.0.0-20210314154223-e6e6c4f2bb5b
	golang.org/x/net v0.0.0-20210924151903-3ad01bbaa167
	golang.org/x/oauth2 v0.0.0-20201208152858-08078c50e5b5
	google.golang.org/api v0.36.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/postgres v1.0.6
	gorm.io/gorm v1.20.11
)

require (
	cloud.google.com/go v0.75.0 // indirect
	github.com/RoaringBitmap/roaring v0.5.5 // indirect
	github.com/antchfx/xpath v1.1.11 // indirect
	github.com/blevesearch/go-porterstemmer v1.0.3 // indirect
	github.com/blevesearch/mmap-go v1.0.2 // indirect
	github.com/blevesearch/segment v0.9.0 // indirect
	github.com/blevesearch/snowballstem v0.9.0 // indirect
	github.com/blevesearch/zap/v11 v11.0.14 // indirect
	github.com/blevesearch/zap/v12 v12.0.14 // indirect
	github.com/blevesearch/zap/v13 v13.0.6 // indirect
	github.com/blevesearch/zap/v14 v14.0.5 // indirect
	github.com/blevesearch/zap/v15 v15.0.3 // indirect
	github.com/cenkalti/hub v1.0.1 // indirect
	github.com/couchbase/ghistogram v0.1.0 // indirect
	github.com/couchbase/moss v0.1.0 // indirect
	github.com/couchbase/vellum v1.0.2 // indirect
	github.com/frankban/quicktest v1.11.3 // indirect
	github.com/glycerine/go-unsnap-stream v0.0.0-20190901134440-81cf024a9e0a // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.2 // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/ishidawataru/sctp v0.0.0-20191218070446-00ab2ac2db07 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.8.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.6.2 // indirect
	github.com/jackc/pgx/v4 v4.10.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.11.12 // indirect
	github.com/lib/pq v1.8.0 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/nats-io/nats-server/v2 v2.2.6 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/steveyen/gtreap v0.1.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tinylib/msgp v1.1.5 // indirect
	github.com/willf/bitset v1.1.11 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.1-0.20180714160509-73f8eece6fdc // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.opencensus.io v0.22.5 // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.6 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210111234610-22ae2b108f89 // indirect
	google.golang.org/grpc v1.34.1 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
