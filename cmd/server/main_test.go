package main

import ("net/http"; "net/http/httptest"; "testing")
func TestHealth(t *testing.T) {
 req:=httptest.NewRequest("GET","/health",nil); rec:=httptest.NewRecorder()
 httpHandler:=http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){ w.WriteHeader(200) })
 httpHandler.ServeHTTP(rec,req); if rec.Code!=200 { t.Fatal(rec.Code) }
}
