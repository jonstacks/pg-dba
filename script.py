from string import ascii_lowercase

def code():
    for a in ascii_lowercase:
        for b in ascii_lowercase:
            for c in ascii_lowercase:
                    yield "".join([a, b, c])
    raise StopIteration


i = 0
for c in code():
    print "INSERT INTO films(code, title, did, date_prod, kind, len)"
    print "VALUES ('{}', 'Title-{}', {}, current_date, 'Kind-{}', random() * interval '3 days');".format(c, i, i, i)
    print ""
    i += 1



# -- Do lots of single inserts to create bloat
# INSERT INTO films(code, title, did, date_prod, kind, len)
# VALUES ('aaaaa', 'Title-1', 10, current_date, 'Kind-1', random() * interval '3 days');
