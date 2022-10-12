package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v4"
	"reflect"
)

var ctx = context.Background()

type Item struct {
	Foo        string
	TimerCache TimerCache
}

type TimerCache struct {
	NasMsg      []byte
	SessRuleIds []string // condition session ruleIds
	PccRuleIds  []string // condition pcc ruleIds
}

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	//marshal, _ := msgpack.Marshal(&Item{Foo: "bar"})
	marshal, _ := HashMarshalValues(&Item{Foo: "bar", TimerCache: TimerCache{
		NasMsg:      []byte("sdf"),
		SessRuleIds: []string{"dsv"},
		PccRuleIds:  []string{"3r4234"},
	}})
	_, err = rdb.HMSet(ctx, "test", marshal).Result()
	if err != nil {
		panic(err)
	}

	var ii Item
	f, err := getFields(&ii)
	fmt.Println("sdfsdf:========:", f)
	if err != nil {
		panic(err)
	}

	vv, err := rdb.HMGet(context.Background(), "test", f...).Result()
	if err != nil {
		panic(err)
	}

	hashUnmarshal(vv, f, &ii)
	fmt.Println("v:", ii)
}
func HashMarshalValues(v interface{}) (map[string]interface{}, error) {
	t, ok := v.(map[string]interface{})
	if ok {
		return t, nil
	}

	m, err := convert(v)
	if err != nil {
		fmt.Println("err:", err.Error())
		return nil, err
	}

	for k, v := range m {
		if !reflect.ValueOf(v).IsNil() {
			buf, err := msgpack.Marshal(v)
			if err != nil {
				return nil, err
			}
			m[k] = buf
		} else {
			delete(m, k)
		}

	}

	return m, nil
}
func convert(v interface{}) (map[string]interface{}, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil, errors.New("reflect.TypeOf fail") // &InvalidHashError{reflect.TypeOf(v)}
	}

	rv = rv.Elem()
	rt := rv.Type()
	m := make(map[string]interface{})

	var tag string
	var ptr interface{}
	for i := 0; i < rt.NumField(); i++ {
		tag = rt.Field(i).Name
		//tag = strings.Trim(rt.Field(i).Tag.Get("redis"), "hash=")

		if rv.Field(i).Kind() != reflect.Ptr {
			ptr = rv.Field(i).Addr().Interface()
		} else {
			ptr = rv.Field(i).Interface()
		}

		if len(tag) > 0 {
			m[tag] = ptr
		}
	}

	return m, nil
}

func HashUnmarshalValues(data map[string]interface{}, v interface{}) error {
	f, err := getFields(v)
	if err != nil {
		return err
	}

	var val []interface{}
	for _, fld := range f {
		a, _ := data[fld]
		val = append(val, a)
	}

	hashUnmarshal(val, f, v)
	return nil
}

func getFields(v interface{}) (f []string, err error) {
	m, err := convert(v)
	if err != nil {
		return
	}

	for k := range m {
		f = append(f, k)
	}

	return
}

func hashUnmarshal(val []interface{}, f []string, v interface{}) {
	rv := reflect.ValueOf(v)
	rt := rv.Type().Elem()
	rv = rv.Elem()

	var ptr interface{}
	for i := 0; i < len(f); i++ {
		if val[i] != nil {
			fl, _ := rt.FieldByName(f[i])
			ft := fl.Type
			fv := rv.FieldByName(f[i])
			if fv.Kind() == reflect.Ptr {
				if fv.IsNil() {
					fv.Set(reflect.New(ft.Elem()))
				}
				ptr = fv.Interface()
			} else {
				ptr = fv.Addr().Interface()
			}

			_ = msgpack.Unmarshal([]byte(val[i].(string)), ptr)
		}
	}
}

func main() {
	ExampleClient()
}
