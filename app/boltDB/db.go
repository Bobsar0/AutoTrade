package app

import (
	"os"
	"github.com/boltdb/bolt/"
)

type BoltDB struct {
	db *bolt.DB
}

func NewBoltDB(filepath string) *BoltDB {
	db, err := bolt.Open(filepath+"/users.db", 0600,  &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	b, err := tx.CreateBucketIfNotExists([]byte("UsersBucket"))
	if err != nil {
		return fmt.Errorf("Error in creating bucket: %s", err)
	}
	return &BoltDB{db}
}

func (b *BoltDB) Close() {
	b.db.Close()
}

func (b *BoltDB) Path() string {
	return b.db.Path()
}

//GetDB starts a read-only transaction
//The Get() function does not return an error because its operation is guaranteed to work 
//Note that values returned from Get() are only valid while the transaction is open. 
//If you need to use a value outside of the transaction then you must use copy() to copy it to another byte slice.
func (b *BoltDB) GetDB(key){
	b.db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("UsersBucket"))
		v := bkt.Get([]byte(key))
		fmt.Printf("Value retrieved is: %s\n", v)
		return nil
	})
}

//DeleteDB deletes a user with a given key from the database
func (b *BoltDB) DeleteDB(key){
	b.db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("UsersBucket"))
		v := bkt.Delete([]byte(key))
		fmt.Printf("Key %s deleted successfully", k)
		return nil
	})
}
// CreateUser saves u to the UsersBucket. The new user ID is set on u once the data is persisted.
func (b *BoltDB) CreateUser(u *model.User) error {
    return b.db.Update(func(tx *bolt.Tx) error {
        // Retrieve the users bucket. This should be created when the DB is first opened.
        bkt := tx.Bucket([]byte("UsersBucket"))

        // Generate ID for the user.
        // This returns an error only if the Tx is closed or not writeable.
        // That can't happen in an Update() call so error check is ignored.
        id, _ := bkt.NextSequence() //lets Bolt determine a sequence which can be used as the unique identifier for the key/value pairs.
        u.ID = int(id)

        // Marshal user data into bytes.
        buf, err := json.Marshal(u)
        if err != nil {
            return err
        }
        // Persist bytes to users bucket.
        return bkt.Put(itob(u.ID), buf)
    })
}

// itob returns an 8-byte big endian representation of v.
// Big endian is an order in which the "big end" (most significant value in the sequence) is stored first (at the lowest storage address)
func itob(v int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}

//Log stats every t seconds:
func (b *BoltDB) logDBStats(t int) {
	// Grab the initial stats.
	prev := b.db.Stats()

	for {
		// Wait for t sec.
		time.Sleep(t * time.Second)

		// Grab the current stats and diff them.
		stats := b.db.Stats()
		diff := stats.Sub(&prev)

		// Encode stats to JSON and print to STDERR.
		json.NewEncoder(os.Stderr).Encode(diff)

		// Save stats for the next loop.
		prev = stats
	}
}()