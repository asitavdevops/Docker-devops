from flask import Flask, render_template, redirect

app = Flask(__name__)

name = "Asitav"
job_profile = "Cloud & DevOps Engineer"

@app.route("/")
def home():
    return render_template("index.html", name=name, job=job_profile)

@app.route("/apply/naukri")
def apply_naukri():
    return redirect("https://www.naukri.com/devops-engineer-jobs")

@app.route("/apply/linkedin")
def apply_linkedin():
    return redirect("https://www.linkedin.com/jobs/devops-engineer-jobs")

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
