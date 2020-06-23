#!/usr/bin/env python3
from http.server import HTTPServer, SimpleHTTPRequestHandler
import sys
import ssl


class CORSRequestHandler (SimpleHTTPRequestHandler):
    def end_headers (self):
        self.send_header('Access-Control-Allow-Origin', '*')
        SimpleHTTPRequestHandler.end_headers(self)


def main():
    host = '0.0.0.0'
    port = 4443
    certfile = 'certs/server.pem'

    httpd = HTTPServer((host, port), CORSRequestHandler)
    httpd.socket = ssl.wrap_socket (httpd.socket, certfile=certfile, server_side=True)

    print('Running server on {0}:{1}'.format(host, port))
    httpd.serve_forever()


if __name__ == '__main__':
    main()
