package mongodemo

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	binary 和 JSON 的合成词, 可以理解为 JSON 文档的二进制表示形式
	BSON 是一种二进制序列化格式, 用于在 MongoDB 中存储文档和进行远程过程调用
*/
/*
When decoding BSON to a D or M, the following type mappings apply when unmarshaling:

BSON int32 unmarshals to an int32.
BSON int64 unmarshals to an int64.
BSON double unmarshals to a float64.
BSON string unmarshals to a string.
BSON boolean unmarshals to a bool.
BSON embedded document unmarshals to the parent type (i.e. D for a D, M for an M).
BSON array unmarshals to a bson.A.
BSON ObjectId unmarshals to a primitive.ObjectID.
BSON datetime unmarshals to a primitive.DateTime.
BSON binary unmarshals to a primitive.Binary.
BSON regular expression unmarshals to a primitive.Regex.
BSON JavaScript unmarshals to a primitive.JavaScript.
BSON code with scope unmarshals to a primitive.CodeWithScope.
BSON timestamp unmarshals to an primitive.Timestamp.
BSON 128-bit decimal unmarshals to an primitive.Decimal128.
BSON min key unmarshals to an primitive.MinKey.
BSON max key unmarshals to an primitive.MaxKey.
BSON undefined unmarshals to a primitive.Undefined.
BSON null unmarshals to nil.
BSON DBPointer unmarshals to a primitive.DBPointer.
BSON symbol unmarshals to a primitive.Symbol.

The above mappings also apply when marshaling a D or M to BSON. Some other useful marshaling mappings are:

time.Time marshals to a BSON datetime.
int8, int16, and int32 marshal to a BSON int32.
int marshals to a BSON int32 if the value is between math.MinInt32 and math.MaxInt32, inclusive, and a BSON int64 otherwise.
int64 marshals to BSON int64 (unless Encoder.IntMinSize is set).
uint8 and uint16 marshal to a BSON int32.
uint, uint32, and uint64 marshal to a BSON int64 (unless Encoder.IntMinSize is set).
BSON null and undefined values will unmarshal into the zero value of a field (e.g. unmarshaling a BSON null or undefined value into a string will yield the empty string.).
*/

type E1 struct {
	Key   string
	Value any
}
type D1 []E1
type D2 = D1

func TestBsonDIY(t *testing.T) {
	t.Run("diy", func(t *testing.T) {
		var d2 D2 = D2{{"foo", "bar"}, {"hello", 3.1415}}
		t.Logf("d2 %T %+v \n", d2, d2)

		var d11 D1 = D1{{"foo", "bar"}, {"hello", 3.1415}}
		t.Logf("d11 %T %+v \n", d11, d11)

		var e11 E1 = E1{"foo", 3.1415}
		t.Logf("e11 %T %+v \n", e11, e11)
	})
}
func TestBSONType(t *testing.T) {
	t.Run("bson.A", func(t *testing.T) {
		// bson.A 是一个有序数组
		var a1 = bson.A{"foo", "bar", "hello", 3.1415}
		t.Logf("type: %T, %+v\n", a1, a1)
	})
	t.Run("bson.D", func(t *testing.T) {
		// bson.D 是一个 struct slice
		// var d1 = bson.D{{"foo", "bar"}, {"hello", 3.1415}}
		// t.Logf("type: %T, %+v\n", d1, d1)

		var d2 = bson.D{{Key: "foo", Value: "bar"}, {Key: "hello", Value: 3.1415}}
		t.Logf("type: %T, %+v\n", d2, d2)
	})
	t.Run("bson.M", func(t *testing.T) {
		// bson.M 是一个map
		var m = bson.M{"foo": "bar", "hello": 3.1415}
		t.Logf("type: %T, %+v\n", m, m)
	})
}
