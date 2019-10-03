from src import app, mysql
from flask import jsonify

grades = {
        'pg': 'perfect_grade', 
        'mg': 'master_grade', 
        'rg': 'real_grade', 
        'hguc': 'hguc'
        }

# TODO: remove
@app.route('/')
@app.route('/index')
def index():
    return "Hello, World!"

@app.route('/api/1.0/<string:grade>', methods=['GET'])
def get_grade(grade):
    # TODO: validate the grade
    # TODO: get list of kits from the given grade
    cursor = mysql.get_db().cursor()

    # TODO: paginate results
    if (grade == 'mg'):
        query = 'select * from master_grade'
        cursor.execute(query)
        data = cursor.fetchall()
    else: 
        # TODO: return a proper error
        return "invalid grade"
        
    result = [{
            'kit_id': kit[1],
            'name': kit[2],
            'series': kit[3],
            'price': kit[4],
            'description': kit[5]
        } for kit in data]

    return jsonify(result) 

@app.route('/api/1.0/<string:grade>/<int:kit_id>', methods=['GET'])
def get_kit(grade, kit_id):
    # TODO: return an error if the grade is invalid
    
    cursor = mysql.get_db().cursor()
    query = "select * from %s where mg_id = %d" % (grades[grade], kit_id)
    cursor.execute(query)
    kit = cursor.fetchone()

    return jsonify({
        'kit_id': kit[1],
        'name': kit[2],
        'series': kit[3],
        'price': kit[4],
        'description': kit[5]
    })

@app.route('/api/1.0/<string:grade>', methods=['POST'])
def insert_kit(grade):
    return grade

@app.route('/api/1.0/<string:grade>/<int:kit_id>', methods=['PUT'])
def update_kit(grade, kit_id):
    return grade, kit_id

@app.route('/api/1.0/<string:grade>/<int:kit_id>', methods=['DELETE'])
def delete_kit(grade, kit_id):
    return grade, kit_id
