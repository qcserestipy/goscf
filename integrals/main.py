# pyscf_service/main.py
from flask import Flask, request, jsonify
from pyscf import gto

app = Flask(__name__)

@app.route('/compute_integrals', methods=['POST'])
def compute_integrals():
    data = request.get_json()
    if not data or 'atom' not in data or 'basis' not in data:
        return jsonify({"error": "Missing required parameters"}), 400

    # Construct the molecule with given atom specification and basis set.
    mol = gto.M(atom=data['atom'], basis=data['basis'])

    # Compute overlap integrals and kinetic energy integrals (as examples)
    s = mol.intor('int1e_ovlp')
    t = mol.intor('int1e_kin')

    # Return integrals as JSON (convert numpy arrays to lists)
    return jsonify({
        "overlap_integrals": s.tolist(),
        "kinetic_integrals": t.tolist()
    })

if __name__ == '__main__':
    app.run(debug=True, port=5000)
