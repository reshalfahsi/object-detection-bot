import json
from flask import Flask, request, jsonify

app = Flask(__name__)
host = ''

class Data:

    def __init__(self, obj={}):
        self._obj = obj

    def get_obj(self):
        return self._obj

    def set_column(self, name, value):
        self._obj[name] = value

    def set_columns(self, columns):
        self._obj.update(columns)

    def delete_column(self, name):
        if name in self._obj:
            del self._obj[name]

    @staticmethod
    def serialize(data):
        return json.dumps(data._obj, sort_keys=True, indent=7)

    @staticmethod
    def deserialize(string):
        return Data(json.loads(string))

@app.route('/', methods=['GET'])
def default():
    return ''

@app.route('/client/<client_id>', methods=['GET'])
def get(client_id):
    data = get_data(client_id)
    if data is None:
        return "Not Available!!"
    else:
        return jsonify(data.get_obj())

@app.route('/client/<client_id>', methods=['POST'])
def post(client_id):
    result = set_column(client_id, json.loads(request.data))
    return jsonify(result.get_obj())

@app.route('/client/<client_id>/<column_name>', methods=['DELETE'])
def delete(client_id, column_name):
    result = delete_column(client_id, column_name)
    if result is None:
        return "Not Available"
    else:
        return jsonify(result.get_obj())

def get_data(client_id):
    return read_data(client_id)

def set_column(client_id, columns):
    data = read_data(client_id)
    if data is None:
        data = Data()
    data.set_columns(columns)
    result = write_data(client_id, data)
    return result

def delete_column(client_id, column_name):
    data = read_data(client_id)
    if data is None:
        return None
    if column_name not in data.get_obj():
        return None
    data.delete_column(column_name)
    result = write_data(client_id, data)
    return result

def read_data(client_id):
    file_name = client_id + '.json'
    try:
        with open(file_name, 'r') as f:
            return Data.deserialize(f.read())
    except IOError as e:
        print e
        return None

def write_data(client_id, data):
    file_name = client_id + '.json'
    try:
        with open(file_name, 'w') as f:
            f.write(Data.serialize(data))
            return data
    except IOError as e:
        print e
        return None

if __name__ == '__main__':
    app.run(debug=True, host=host)
