package utils

// Blogs Table Query
const (
	LIST_QUERY   = `SELECT * FROM blogs ORDER BY published_at DESC`
	GET_QUERY    = `SELECT id, user_id, title, content, published_at, archive, comment_count FROM blogs WHERE id=?`
	POST_QUERY   = `INSERT INTO blogs (id, user_id, title, content,archive,comment_count) VALUES (?,?,?,?,?,?)`
	UPDATE_QUERY = `UPDATE blogs SET title=?,content=?,published_at=?, comment_count=? WHERE id=?`
	DELETE_QUERY = `DELETE FROM blogs WHERE id=?`
)

// Get hashed password
const (
	GET_HASHED_QUERY = `SELECT id, firstname, lastname, username, hashed_password FROM users WHERE email=?`

	COMPARE_HASH_ERR        = `Incorrect password.`
	COMPARE_HASH_ERR_DETAIL = `You entered an incorrect password.`
)

// Links Table Query
const (
	LINK_INSERT_QUERY = `INSERT INTO links (id,blog_id,rel,href) VALUES (?,?,?,?),(?,?,?,?)`
)

// User registration query
const (
	INSERT_USER_QUERY = `INSERT INTO users (id, first_name,last_name,username,email,hashed_password,hashed_passkey) VALUES (?,?,?,?,?,?,?)`
	CHECK_USER_QUERY  = `SELECT EXISTS(SELECT 1 FROM users WHERE email=? OR username=?)`
)

// Error Strings
const (
	BAD_REQ_ERROR         = `Bad Request`
	BAD_REQ_ERROR_DETAILS = `Invalid blog format.`
	BAD_REQ_ERROR_CODE    = `400`

	CONN_ERR = `Connection Error.`

	NOT_FOUND_ERR        = `Not Found.`
	NOT_FOUND_ERR_CODE   = `404`
	NOT_FOUND_ERR_DETAIL = `Blog was not found on the database.`

	SERVER_ERROR        = `Internal server error.`
	SERVER_ERROR_CODE   = `500`
	SERVER_ERROR_DETAIL = `Unable to connect to Server.`

	DELETE_ERROR           = `Deletion error.`
	ROW_SCAN_ERR           = `Row scan error`
	INSERTION_ERR          = `Insertion error.`
	ZERO_ROWS_AFFECTED_ERR = `No rows affected.`
	SAME_TITLE             = `The title is the same as before.`
	SAME_CONTENT           = `The content is the same as before.`

	ALREADY_EXISTS_ERROR        = `Already exists.`
	ALREADY_EXISTS_ERROR_DETAIL = `The user already exists. (Hint: Kindly check your email or username.)`
	ALREADY_EXISTS_ERROR_CODE   = `409`
)

// Bcrypt hash password
const (
	LONG_PASS_ERR        = `Password is too long.`
	LONG_PASS_ERR_DETAIL = `The password you entered is too long, consider shortening it but not making it too simple.`

	SHORT_PASS_ERR        = `Password is too short.`
	SHORT_PASS_ERR_DETAIL = `The password you entered is too short. Password should be 8 characters long.`

	PASS_MISMATCH_ERR        = `Password mismatch.`
	PASS_MISMATCH_ERR_DETAIL = `The password you entered is incorrect.`

	INVALID_PASSWORD_ERROR        = `Invalid password`
	INVALID_PASSWORD_ERROR_DETAIL = `The password entered contains invalid characters.`
)

// User Input validations
const (
	EMAIL_EMPTY     = `Email should not be empty.`
	EMPTY_FIELD     = `Empty field.`
	FIRSTNAME_EMPTY = `Firstname should not be empty.`
	LASTNAME_EMPTY  = `Lastname should not be empty.`
	PASSWORD_EMPTY  = `Password should not be empty.`
	SHORTPASSWORD   = `Password should should be at least 6 characters.`
	PASSKEY_EMPTY   = `Passkey should not be empty.`
	LONGPASSKEY     = `Passkey should be at most 6 characters.`
	USERNAME_EMPTY  = `Username should not be empty.`
	SUCCESS_MESSAGE = `You have been registered successfully.`
)
