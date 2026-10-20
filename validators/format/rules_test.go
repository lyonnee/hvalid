package format

import (
	"strings"
	"testing"
)

func requireErr(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
	if err.Error() != want {
		t.Fatalf("got %q, want %q", err.Error(), want)
	}
}

func requireNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUUID(t *testing.T) {
	v4 := "550e8400-e29b-41d4-a716-446655440000"
	v3 := "550e8400-e29b-31d4-a716-446655440000"
	requireNil(t, UUID[string]()(v4))
	requireNil(t, UUID[string]()(strings.ToUpper(v4)))
	requireNil(t, UUID[string](3, 4)(v4))
	requireNil(t, UUID[string](4)(v4))
	requireErr(t, UUID[string](3)(v4), "must be a UUID")
	requireErr(t, UUID[string]()(v3+"x"), "must be a UUID")
	requireErr(t, UUID[string]()("urn:uuid:"+v4), "must be a UUID")
	requireErr(t, UUID[string]()("{"+v4+"}"), "must be a UUID")
	requireErr(t, UUID[string]()("550e8400-e29b-41d4-1716-446655440000"), "must be a UUID")
	requireNil(t, UUID[string]()(v3))
}

func TestULID(t *testing.T) {
	id := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
	requireNil(t, ULID[string]()(id))
	requireNil(t, ULID[string]()(strings.ToLower(id)))
	requireErr(t, ULID[string]()(id[:25]), "must be a ULID")
	requireErr(t, ULID[string]()("01ARZ3NDEKTSV4RRFFQ69G5FAI"), "must be a ULID")
}

func TestJSON(t *testing.T) {
	requireNil(t, JSON[string]()("null"))
	requireNil(t, JSON[string]()("1"))
	requireNil(t, JSON[string]()("{}"))
	requireNil(t, JSON[string]()("[]"))
	requireErr(t, JSON[string]()(""), "must be JSON")
	requireErr(t, JSON[string]()("{"), "must be JSON")
}

func TestBase64(t *testing.T) {
	requireNil(t, Base64[string]()("YQ=="))
	requireErr(t, Base64[string]()(""), "must be base64")
	requireErr(t, Base64[string]()("YQ==\n"), "must be base64")
	requireErr(t, Base64[string]()("YQ"), "must be base64")
	requireErr(t, Base64[string]()("****"), "must be base64")
}

func TestDataURI(t *testing.T) {
	requireNil(t, DataURI[string]()("data:,hello"))
	requireNil(t, DataURI[string]()("DATA:,hello"))
	requireNil(t, DataURI[string]()("data:text/plain;base64,YQ=="))
	requireNil(t, DataURI[string]()("data:text/plain;BASE64,YQ=="))
	requireErr(t, DataURI[string]()(""), "must be a data URI")
	requireErr(t, DataURI[string]()("http:"), "must be a data URI")
	requireErr(t, DataURI[string]()("data:hello"), "must be a data URI")
	requireErr(t, DataURI[string]()("data:text/plain;base64,%%%"), "must be a data URI")
}

func TestMAC(t *testing.T) {
	requireNil(t, MAC[string]()("00:11:22:33:44:55"))
	requireNil(t, MAC[string]()("00-11-22-33-44-55"))
	requireErr(t, MAC[string]()(""), "must be a MAC address")
	requireErr(t, MAC[string]()("zz"), "must be a MAC address")
}

func TestDNSName(t *testing.T) {
	requireNil(t, DNSName[string]()("localhost"))
	requireNil(t, DNSName[string]()("Example.COM"))
	requireNil(t, DNSName[string]()("a1.com"))
	requireNil(t, DNSName[string]()(strings.Repeat("a", 63)))
	exact := strings.Repeat("a", 63) + "." + strings.Repeat("b", 63) + "." + strings.Repeat("c", 63) + "." + strings.Repeat("d", 61)
	requireNil(t, DNSName[string]()(exact))

	requireErr(t, DNSName[string]()(""), "must be a DNS name")
	requireErr(t, DNSName[string]()("example.com."), "must be a DNS name")
	requireErr(t, DNSName[string]()(exact+"e"), "must be a DNS name")
	requireErr(t, DNSName[string]()(strings.Repeat("a", 64)), "must be a DNS name")
	requireErr(t, DNSName[string]()("-a.com"), "must be a DNS name")
	requireErr(t, DNSName[string]()("a-.com"), "must be a DNS name")
	requireErr(t, DNSName[string]()("a..b"), "must be a DNS name")
	requireErr(t, DNSName[string]()("a_b"), "must be a DNS name")
}

func TestHost(t *testing.T) {
	requireNil(t, Host[string]()("127.0.0.1"))
	requireNil(t, Host[string]()("::1"))
	requireNil(t, Host[string]()("example.com"))
	requireErr(t, Host[string]()("[::1]"), "must be a host")
	requireErr(t, Host[string]()(""), "must be a host")
	requireErr(t, Host[string]()("not host"), "must be a host")
}

func TestPort(t *testing.T) {
	requireNil(t, Port[string]()("0"))
	requireNil(t, Port[string]()("65535"))
	requireNil(t, Port[string]()("00080"))
	requireErr(t, Port[string]()("65536"), "must be a port")
	requireErr(t, Port[string]()(""), "must be a port")
	requireErr(t, Port[string]()("-1"), "must be a port")
	requireErr(t, Port[string]()("+80"), "must be a port")
	requireErr(t, Port[string]()("80a"), "must be a port")
}

func TestMongoID(t *testing.T) {
	requireNil(t, MongoID[string]()("507f1f77bcf86cd799439011"))
	requireNil(t, MongoID[string]()("507f1f77bcf86cd7994390AB"))
	requireErr(t, MongoID[string]()("507f1f77bcf86cd79943901"), "must be a MongoDB object id")
	requireErr(t, MongoID[string]()("507f1f77bcf86cd79943901g"), "must be a MongoDB object id")
}

func TestJWT(t *testing.T) {
	header := "eyJhbGciOiJub25lIn0"
	payload := "eyJzdWIiOiJhIn0"
	requireNil(t, JWT[string]()(header+"."+payload+".c2ln"))
	requireNil(t, JWT[string]()(header+"."+payload+".YQ=="))
	requireErr(t, JWT[string]()(header+"."+payload), "must be a JWT")
	requireErr(t, JWT[string]()("@@@."+payload+".c2ln"), "must be a JWT")
	requireErr(t, JWT[string]()(header+".YQ.c2ln"), "must be a JWT")
	requireErr(t, JWT[string]()(header+".bnVsbA.c2ln"), "must be a JWT")
	requireErr(t, JWT[string]()(header+".W10.c2ln"), "must be a JWT")
	requireErr(t, JWT[string]()(header+"."+payload+".@@@"), "must be a JWT")
	requireErr(t, JWT[string]()(header+"..c2ln"), "must be a JWT")
	requireErr(t, JWT[string]()("+++."+payload+".c2ln"), "must be a JWT")
	requireErr(t, JWT[string]()(header+"."+payload+".===="), "must be a JWT")
}

func TestSemver(t *testing.T) {
	requireNil(t, Semver[string]()("1.2.3"))
	requireNil(t, Semver[string]()("1.2.3-alpha.1+build"))
	requireErr(t, Semver[string]()("1.2"), "must be a semantic version")
	requireErr(t, Semver[string]()("01.2.3"), "must be a semantic version")
	requireErr(t, Semver[string]()("v1.2.3"), "must be a semantic version")
	requireErr(t, Semver[string]()(""), "must be a semantic version")
}

func TestColors(t *testing.T) {
	requireNil(t, HexColor[string]()("#fff"))
	requireNil(t, HexColor[string]()("#336699"))
	requireNil(t, HexColor[string]()("#336699ff"))
	requireNil(t, HexColor[string]()("#FFF"))
	requireErr(t, HexColor[string]()("#gg0000"), "must be a hex color")
	requireErr(t, HexColor[string]()("336699"), "must be a hex color")
	requireErr(t, HexColor[string]()("#ffff"), "must be a hex color")
	requireErr(t, HexColor[string]()(""), "must be a hex color")

	requireNil(t, RGB[string]()("rgb(0, 0, 0)"))
	requireNil(t, RGB[string]()("rgb(0,0,0)"))
	requireNil(t, RGB[string]()("rgb(255, 255, 255)"))
	requireErr(t, RGB[string]()("rgb(256, 0, 0)"), "must be an RGB color")
	requireErr(t, RGB[string]()("rgba(0, 0, 0, 1)"), "must be an RGB color")
	requireErr(t, RGB[string]()(""), "must be an RGB color")
}

func TestCoordinates(t *testing.T) {
	requireNil(t, Latitude[string]()("90"))
	requireNil(t, Latitude[string]()("-90"))
	requireNil(t, Latitude[string]()("0"))
	requireErr(t, Latitude[string]()("90.1"), "must be a latitude")
	requireErr(t, Latitude[string]()("-90.1"), "must be a latitude")
	requireErr(t, Latitude[string]()(""), "must be a latitude")
	requireErr(t, Latitude[string]()("abc"), "must be a latitude")
	requireErr(t, Latitude[string]()("NaN"), "must be a latitude")
	requireErr(t, Latitude[string]()("Inf"), "must be a latitude")

	requireNil(t, Longitude[string]()("180"))
	requireNil(t, Longitude[string]()("-180"))
	requireNil(t, Longitude[string]()("0"))
	requireErr(t, Longitude[string]()("180.1"), "must be a longitude")
	requireErr(t, Longitude[string]()("-180.1"), "must be a longitude")
}

func TestTimeFormats(t *testing.T) {
	requireNil(t, RFC3339[string]()("2026-10-19T16:42:00Z"))
	requireNil(t, RFC3339[string]()("2026-10-19T16:42:00.5Z"))
	requireErr(t, RFC3339[string]()("2026-10-19"), "must be RFC3339")

	requireNil(t, UnixTime[string]()("0"))
	requireNil(t, UnixTime[string]()("-1"))
	requireNil(t, UnixTime[string]()("00"))
	requireNil(t, UnixTime[string]()("9223372036854775807"))
	requireNil(t, UnixTime[string]()("-9223372036854775808"))
	requireErr(t, UnixTime[string]()(""), "must be a unix time")
	requireErr(t, UnixTime[string]()("-"), "must be a unix time")
	requireErr(t, UnixTime[string]()("+1"), "must be a unix time")
	requireErr(t, UnixTime[string]()("1.5"), "must be a unix time")
	requireErr(t, UnixTime[string]()("9223372036854775808"), "must be a unix time")

	requireNil(t, YYYYMMDD[string]()("20261019"))
	requireErr(t, YYYYMMDD[string]()("20260231"), "must be a YYYYMMDD date")
	requireErr(t, YYYYMMDD[string]()("20261301"), "must be a YYYYMMDD date")
	requireErr(t, YYYYMMDD[string]()("2026101"), "must be a YYYYMMDD date")
	requireErr(t, YYYYMMDD[string]()("2026101a"), "must be a YYYYMMDD date")
	requireErr(t, YYYYMMDD[string]()(""), "must be a YYYYMMDD date")
}

func TestISBN(t *testing.T) {
	requireNil(t, ISBN[string]()("0-306-40615-2"))
	requireNil(t, ISBN[string]()("978-0-306-40615-7"))
	requireNil(t, ISBN[string]()("978 0 306 40615 7"))
	requireNil(t, ISBN[string]()("0-8044-2957-X"))
	requireNil(t, ISBN[string]()("080442957x"))
	requireNil(t, ISBN[string]()("9790000000001"))
	requireErr(t, ISBN[string]()(""), "must be an ISBN")
	requireErr(t, ISBN[string]()("0306406153"), "must be an ISBN")
	requireErr(t, ISBN[string]()("0306406A52"), "must be an ISBN")
	requireErr(t, ISBN[string]()("X306406152"), "must be an ISBN")
	requireErr(t, ISBN[string]()("9770306406157"), "must be an ISBN")
	requireErr(t, ISBN[string]()("978030640615X"), "must be an ISBN")
	requireErr(t, ISBN[string]()("9780306406158"), "must be an ISBN")
	requireErr(t, ISBN[string]()("---"), "must be an ISBN")
}

func TestISOCodes(t *testing.T) {
	requireNil(t, ISO3166Alpha2[string]()("CN"))
	requireNil(t, ISO3166Alpha2[string]()("us"))
	requireErr(t, ISO3166Alpha2[string]()("XX"), "must be an ISO 3166-1 alpha-2 code")
	requireErr(t, ISO3166Alpha2[string]()("USA"), "must be an ISO 3166-1 alpha-2 code")
	requireErr(t, ISO3166Alpha2[string]()(""), "must be an ISO 3166-1 alpha-2 code")

	requireNil(t, ISO4217[string]()("CNY"))
	requireNil(t, ISO4217[string]()("usd"))
	requireErr(t, ISO4217[string]()("US"), "must be an ISO 4217 code")
	requireErr(t, ISO4217[string]()(""), "must be an ISO 4217 code")
}
