from flask import Flask, request, jsonify
import os
import json
from datetime import datetime

app = Flask(__name__)

@app.route('/v1/log', methods=['POST'])
def log_endpoint():
    if request.method == 'POST':
        if request.is_json:
            json_data = request.get_json()

            if not os.path.exists('logs'):
                os.makedirs('logs')

            current_month_year = datetime.now().strftime('%m-%Y')
            log_dir = os.path.join('logs', current_month_year)

            if not os.path.exists(log_dir):
                os.makedirs(log_dir)

            log_filename = datetime.now().strftime('%d-%m-%Y-access-log-domain-%H.log')
            log_path = os.path.join(log_dir, log_filename)

            with open(log_path, 'a') as log_file:
                log_file.write(json.dumps(json_data) + '\n')

            return jsonify({'message': 'Log registrado com sucesso!'}), 200
        else:
            return jsonify({'error': 'Conteúdo da requisição não é um JSON'}), 400
    else:
        return jsonify({'error': 'Método não permitido'}), 405

if __name__ == '__main__':
    app.run(port=9832)
