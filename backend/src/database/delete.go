package database

func DeleteImage(imageId, userId int) error {
	db, err := createConnection()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM public.image WHERE (pk = $1) AND (user_id = $2);", imageId, userId)
	if err != nil {
		return err
	}
	return nil
}
