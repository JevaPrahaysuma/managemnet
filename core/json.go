package core

import (
	"fmt"
	"time"
)

// JSONTime is type of time
type JSONTime time.Time

// MarshalJSON is used for marshalling in json
func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	tx := time.Time(t)
	// if tx.IsZero() {
	// 	return nil, nil
	// }
	stamp := fmt.Sprintf("\"%s\"", tx.Format("15:04:05"))
	return []byte(stamp), nil
}

// JSONDate is type of date
type JSONDate time.Time

// MarshalJSON is used for marshalling in json
func (t JSONDate) MarshalJSON() ([]byte, error) {
	//do your serializing here
	tx := time.Time(t)
	// if tx.IsZero() {
	// 	return nil, nil
	// }
	stamp := fmt.Sprintf("\"%s\"", tx.Format("2006-01-02"))
	return []byte(stamp), nil
}

// JSONDateTime is type of datetime
type JSONDateTime time.Time

// MarshalJSON is used for marshalling in json
func (t JSONDateTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	tx := time.Time(t)
	temp1 := tx.Format(time.RFC3339)
	// fmt.Println(tx)
	// if tx.IsZero() {
	// 	return nil, nil
	// }
	temp, _ := time.Parse("2006-01-02 15:04:05", temp1)
	//fmt.Println(temp)
	stamp := fmt.Sprintf("\"%s\"", temp)
	return []byte(stamp), nil
}

// JSONDateTime is type of datetime
type JSONFirestoreTimestamp time.Time

// MarshalJSON is used for marshalling in json
func (t JSONFirestoreTimestamp) MarshalJSON() ([]byte, error) {
	//do your serializing here
	tx := time.Time(t)
	// if tx.IsZero() {
	// 	return nil, nil
	// }
	stamp := fmt.Sprintf("\"%s\"", tx.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
