import cotu

def test_hello_world(capsys):
    cotu.hello_world()
    out, err = capsys.readouterr()
    
    assert out == 'hello world\n'
    assert err == ''
