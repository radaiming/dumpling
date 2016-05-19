#!/usr/bin/env python3
# Created @ 2016-05-20 01:18 by @radaiming
#
# tests like https://github.com/bahlo/goat/blob/master/middleware_test.go
# is much less meaningful, so I run this script to test with my sample programs
# and check their output
#


import subprocess
import os
import shlex
import sys


test_url = 'http://127.0.0.1:9988/'


def test(binary, curl_and_output_checker, check_side='client'):
    # check if what we get is what we expected
    # by default we check client side(curl) output
    subprocess.check_call(('go build samples/%s.go' % binary).split())
    p = subprocess.Popen(['./' + binary], stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
    # ensure process started and print 'now serving on 127.0.0.1:9988'
    p.stdout.readline()

    failed = False
    try:
        for curl_cmd, checker in curl_and_output_checker.items():
            if check_side == 'client':
                output = subprocess.check_output(shlex.split(curl_cmd), stderr=subprocess.STDOUT)
            elif check_side == 'server':
                subprocess.check_call(shlex.split(curl_cmd), stdout=subprocess.PIPE)
                output = p.stdout.readline()
            print('-' * 70)
            print(output.decode('utf-8'))
            print('-' * 70)
            if not checker(output):
                failed = True
                print('\033[91mtest failed for %s, %s\033[0m' % (binary, check_side))
    finally:
        p.terminate()
        os.remove(binary)
    if failed:
        sys.exit(1)
    print('\033[92mtest OK for %s.go\033[0m' % binary)


def test_basic_auth():
    curl_and_output_checker = {
        'curl -s -H "Authorization: Basic dXNlcjAwMTp0b2tlbjAwMQ==" ' + test_url:
            lambda x: x == b'Top Secret Content!'
    }
    test('basic_auth', curl_and_output_checker)


def test_hello():
    curl_and_output_checker = {
        'curl -s ' + test_url: lambda x: x == b'hello world',
        'curl -X POST -s ' + test_url: lambda x: x == b'hello world',
    }
    test('hello', curl_and_output_checker)


def test_logging_demo():
    curl_and_output_checker = {
        'curl -s -H "User-Agent: blabla" ' + test_url: lambda x: b'blabla' in x
    }
    test('logging_demo', curl_and_output_checker, 'server')


def test_middleware():
    curl_and_output_checker = {
        'curl -s ' + test_url: lambda x: b'middleware appended' in x
    }
    test('middleware', curl_and_output_checker)


def test_process_time_logger():
    curl_and_output_checker = {
        'curl -s ' + test_url: lambda x: b' ms' in x
    }
    test('process_time_logger_demo', curl_and_output_checker, 'server')


def test_redirect():
    curl_and_output_checker = {
        'curl -D - ' + test_url: lambda x: b'Location: https://google.com' in x
    }
    test('redirect', curl_and_output_checker)


def test_regex():
    curl_and_output_checker = {
        'curl -s ' + test_url + '/xxx/blabla/': lambda x: x == b'URL matches!',
        'curl -s ' + test_url + '/blabla/': lambda x: not x
    }
    test('regex', curl_and_output_checker)


def test_return_forms():
    curl_and_output_checker = {
        'curl -s -d "a=1&b=2&c=3" ' + test_url: lambda x: b'b -> 2' in x
    }
    test('return_forms', curl_and_output_checker)


def test_return_hash():
    curl_and_output_checker = {
        'curl -s -F 1="xxx" ' + test_url: lambda x: b'b60d121b438a380c343d5ec3c2037564b82ffef3' in x
    }
    test('return_hash', curl_and_output_checker)


def test_return_query():
    curl_and_output_checker = {
        'curl -s ' + test_url + '?a=1&b=2&c=3': lambda x: b'a -> 1' in x
    }
    test('return_query', curl_and_output_checker)


def test_static():
    tmp_file = '/tmp/b60d121'
    with open(tmp_file, 'w') as fd:
        fd.write('7564b82ffef3')
    try:
        curl_and_output_checker = {
            'curl -s ' + test_url + 'static/b60d121': lambda x: b'7564b82ffef3' in x
        }
        test('static', curl_and_output_checker)
    finally:
        os.remove(tmp_file)


def main():
    test_basic_auth()
    test_hello()
    test_logging_demo()
    test_middleware()
    test_process_time_logger()
    test_redirect()
    test_regex()
    test_return_forms()
    test_return_hash()
    test_return_query()
    test_static()


if __name__ == '__main__':
    main()
