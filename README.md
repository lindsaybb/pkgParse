# pkgParse

pkgParse is a generic Golang utility program for summarizing the imports and functions of a given package.

| Flag | Description |
| ------ | ------ |
| -h | Show this help |
| -d | Directory to search for files (default ".") |
| -t | File suffix to parse (default 'go', others unoptimized) |
| -r | Recurse into sub-directories (default false) |
| -o | Output file name (generated from package if blank) |
| -so | Dump output to stdout instead of file |

# Example Input
```sh
./pkgParse -d ~/go/src/github.com/secsy/goftp/ -so
```

# Example Output
```sh
Filename: /home/user/go/src/github.com/secsy/goftp/client.go    Packagename: goftp      Length: 442
Imports: ["crypto/tls" "errors" "fmt" "io" "net" "sync" "time"]
Functions:
        func (e ftpError) Error() string {
        func (e ftpError) Temporary() bool {
        func (e ftpError) Timeout() bool {
        func (e ftpError) Code() int {
        func (e ftpError) Message() string {
        func newClient(config Config, hosts []string) *Client {
        func (c *Client) Close() error {
        func (c *Client) debug(f string, args ...interface{}) {
        func (c *Client) numOpenConns() int {
        func (c *Client) getIdleConn() (*persistentConn, error) {
        func (c *Client) removeConn(pconn *persistentConn) {
        func (c *Client) returnConn(pconn *persistentConn) {
        func (c *Client) OpenRawConn() (RawConn, error) {
        func (c *Client) openConn(idx int, host string) (pconn *persistentConn, err error) {

Filename: /home/user/go/src/github.com/secsy/goftp/client_test.go       Packagename: goftp      Length: 150
Imports: ["bytes" "crypto/tls" "sync" "testing" "time"]
Functions:
        func TestTimeoutConnect(t *testing.T) {
        func TestExplicitTLS(t *testing.T) {
        func TestImplicitTLS(t *testing.T) {
        func TestPooling(t *testing.T) {

Filename: /home/user/go/src/github.com/secsy/goftp/examples_test.go     Packagename: goftp_test Length: 97
Imports: ["bytes" "fmt" "io/ioutil" "os" "time"  "github.com/secsy/goftp"]
Functions:
        func Example() {
        func Example_config() {
        func ExampleClient_OpenRawConn() {

Filename: /home/user/go/src/github.com/secsy/goftp/file_system.go       Packagename: goftp      Length: 496
Imports: ["bufio" "fmt" "os" "path/filepath" "regexp" "strconv" "strings" "time"]
Functions:
        func (c *Client) Delete(path string) error {
        func (c *Client) Rename(from, to string) error {
        func (c *Client) Mkdir(path string) (string, error) {
        func (c *Client) Rmdir(path string) error {
        func (c *Client) Getwd() (string, error) {
        func commandNotSupporterdError(err error) bool {
        func (c *Client) ReadDir(path string) ([]os.FileInfo, error) {
        func (c *Client) Stat(path string) (os.FileInfo, error) {
        func extractDirName(msg string) (string, error) {
        func (c *Client) controlStringList(f string, args ...interface{}) ([]string, error) {
        func (c *Client) dataStringList(f string, args ...interface{}) ([]string, error) {
        func (f *ftpFile) Name() string {
        func (f *ftpFile) Size() int64 {
        func (f *ftpFile) Mode() os.FileMode {
        func (f *ftpFile) ModTime() time.Time {
        func (f *ftpFile) IsDir() bool {
        func (f *ftpFile) Sys() interface{} {
        func parseLIST(entry string, loc *time.Location, skipSelfParent bool) (os.FileInfo, error) {
        func parseMLST(entry string, skipSelfParent bool) (os.FileInfo, error) {

Filename: /home/user/go/src/github.com/secsy/goftp/file_system_test.go  Packagename: goftp      Length: 482
Imports: ["bytes" "fmt" "io/ioutil" "os" "path" "reflect" "sort" "testing" "time"]
Functions:
        func TestDelete(t *testing.T) {
        func TestRename(t *testing.T) {
        func TestMkdirRmdir(t *testing.T) {
        func mustParseTime(f, s string) time.Time {
        func TestParseMLST(t *testing.T) {
        func compareFileInfos(a, b os.FileInfo) error {
        func TestReadDir(t *testing.T) {
        func TestReadDirNoMLSD(t *testing.T) {
        func TestStat(t *testing.T) {
        func TestStatNoMLST(t *testing.T) {
        func TestGetwd(t *testing.T) {

Filename: /home/user/go/src/github.com/secsy/goftp/goftp.go     Packagename: goftp      Length: 89
Imports: ["errors" "fmt" "net" "regexp"]
Functions:
        func Dial(hosts ...string) (*Client, error) {
        func DialConfig(config Config, hosts ...string) (*Client, error) {
        func lookupHosts(hosts []string, ipv6Lookup bool) ([]string, error) {

Filename: /home/user/go/src/github.com/secsy/goftp/main_test.go Packagename: goftp      Length: 148
Imports: ["errors" "fmt" "log" "net" "os" "os/exec" "path" "testing" "time"]
Functions:
        func TestMain(m *testing.M) {
        func startPureFTPD(addrs []string, binary string) (func(), error) {
        func startProFTPD() (func(), error) {

Filename: /home/user/go/src/github.com/secsy/goftp/parallel_walk_test.go        Packagename: goftp_test Length: 84
Imports: ["fmt" "os" "path" "path/filepath" "sync/atomic"  "github.com/secsy/goftp"]
Functions:
        func ExampleClient_ReadDir_parallelWalk() {
        func Walk(client *goftp.Client, root string, walkFn filepath.WalkFunc) (ret error) {

Filename: /home/user/go/src/github.com/secsy/goftp/persistent_connection.go     Packagename: goftp      Length: 546
Imports: ["bufio" "crypto/tls" "fmt" "net" "net/textproto" "strconv" "strings" "time"]
Functions:
        func (pconn *persistentConn) SendCommand(f string, args ...interface{}) (int, string, error) {
        func (pconn *persistentConn) PrepareDataConn() (func() (net.Conn, error), error) {
        func (pconn *persistentConn) ReadResponse() (int, string, error) {
        func (pconn *persistentConn) Close() error {
        func (pconn *persistentConn) setControlConn(conn net.Conn) {
        func (pconn *persistentConn) close() error {
        func (pconn *persistentConn) sendCommandExpected(expected int, f string, args ...interface{}) error {
        func (pconn *persistentConn) sendCommand(f string, args ...interface{}) (int, string, error) {
        func (pconn *persistentConn) readResponse() (int, string, error) {
        func (pconn *persistentConn) debug(f string, args ...interface{}) {
        func (pconn *persistentConn) fetchFeatures() error {
        func (pconn *persistentConn) hasFeature(name string) bool {
        func (pconn *persistentConn) hasFeatureWithArg(name, arg string) bool {
        func (pconn *persistentConn) logIn() error {
        func (pconn *persistentConn) requestPassive() (string, error) {
        func (c *dataConn) Read(buf []byte) (int, error) {
        func (c *dataConn) Write(buf []byte) (int, error) {
        func (pconn *persistentConn) prepareDataConn() (func() (net.Conn, error), error) {
        func (pconn *persistentConn) listenActive() (*net.TCPListener, error) {
        func (pconn *persistentConn) setType(t string) error {
        func (pconn *persistentConn) logInTLS() error {

Filename: /home/user/go/src/github.com/secsy/goftp/raw_conn_test.go     Packagename: goftp      Length: 77
Imports: ["io/ioutil" "strings" "testing"]
Functions:
        func TestRawConn(t *testing.T) {

Filename: /home/user/go/src/github.com/secsy/goftp/reply_codes.go       Packagename: goftp      Length: 75
Imports: []
Functions:
        func positiveCompletionReply(code int) bool {
        func positivePreliminaryReply(code int) bool {
        func transientNegativeCompletionReply(code int) bool {

Filename: /home/user/go/src/github.com/secsy/goftp/transfer.go  Packagename: goftp      Length: 269
Imports: ["fmt" "io" "os" "strconv"]
Functions:
        func (c *Client) Retrieve(path string, dest io.Writer) error {
        func (c *Client) Store(path string, src io.Reader) error {
        func (c *Client) transferFromOffset(path string, dest io.Writer, src io.Reader, offset int64) (int64, error) {
        func (c *Client) size(path string) (int64, error) {
        func (c *Client) canResume() bool {

Filename: /home/user/go/src/github.com/secsy/goftp/transfer_test.go     Packagename: goftp      Length: 432
Imports: ["bytes" "errors" "io/ioutil" "math/rand" "os" "reflect" "strings" "testing" "time"]
Functions:
        func TestRetrieve(t *testing.T) {
        func TestRetrievePASV(t *testing.T) {
        func TestRetrieveActive(t *testing.T) {
        func (tb *testWriter) Write(p []byte) (int, error) {
        func TestResumeRetrieveOnWriteError(t *testing.T) {
        func TestResumeRetrieveOnReadError(t *testing.T) {
        func TestStore(t *testing.T) {
        func TestStoreActive(t *testing.T) {
        func TestStoreError(t *testing.T) {
        func (ts *testSeeker) Read(p []byte) (int, error) {
        func (ts *testSeeker) Seek(offset int64, whence int) (int64, error) {
        func randomBytes(b []byte) {
        func TestResumeStoreOnWriteError(t *testing.T) {
        func TestEmptyLinesFeat(t *testing.T) {

```
