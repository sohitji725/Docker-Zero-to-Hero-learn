FROM ubuntu

WORKDIR /app

COPY requirements.txt /app/
COPY devops /app/devops/

RUN apt-get update && \
    apt-get install -y python3 python3-pip && \
    pip install -r requirements.txt

WORKDIR /app/devops

EXPOSE 8000

ENTRYPOINT ["python3"]
CMD ["manage.py", "runserver", "0.0.0.0:8000"]  


## 🧱 Step-by-step explanation of your Dockerfile

### 1. `FROM ubuntu`

This is your **base image**.
Every Docker image starts from another image (like Ubuntu, Alpine, or Python).
It’s like saying: “Start with an Ubuntu OS inside my container.”

👉 It gives you a minimal Linux environment to work with.

---

### 2. `WORKDIR /app`

This sets the **working directory** inside the container.
Every command you run after this happens inside `/app`.
It’s like doing:

```bash
cd /app
```

inside the container.

If the folder doesn’t exist, Docker creates it automatically.

---

### 3. `COPY requirements.txt /app/`

This copies your **requirements.txt** (dependency list) from your local machine into the container’s `/app` folder.

Now the container has your dependency list ready to install.

---

### 4. `COPY devops /app/devops/`

This copies your **application code** (the `devops` folder) into the container.
So your Django project files (like `manage.py`, `settings.py`, etc.) are now available inside `/app/devops`.

---

### 5. `RUN apt-get update && apt-get install -y python3 python3-pip && pip install -r requirements.txt`

This line does three things in one go:

1. Updates Ubuntu’s package list (`apt-get update`)
2. Installs Python and pip (`apt-get install -y python3 python3-pip`)
3. Uses pip to install all dependencies from `requirements.txt`

This ensures your container has everything Django needs to run.

---

### 6. `WORKDIR /app/devops`

After installing dependencies, you switch the working directory again — this time to your project folder where `manage.py` lives.

That way, when the container starts, it runs commands from the right place.

---

### 7. `EXPOSE 8000`

This tells Docker that your app will use **port 8000** inside the container.
It doesn’t open the port by itself but serves as documentation (and helps tools like Docker Compose map ports properly).

---

### 8. `ENTRYPOINT ["python3"]`

This sets the **main executable** for your container.
It means every time the container starts, it will run:

```bash
python3 <whatever comes next>
```

---

### 9. `CMD ["manage.py", "runserver", "0.0.0.0:8000"]`

This provides **default arguments** to the `ENTRYPOINT`.
Combined together, it runs:

```bash
python3 manage.py runserver 0.0.0.0:8000
```

which starts your Django server.

---

## 🧠 Summary

When you run:

```bash
docker run -p 8000:8000 my-django-app
```

Docker does this behind the scenes:

1. Starts an Ubuntu container
2. Installs Python and dependencies
3. Moves to `/app/devops`
4. Runs `python3 manage.py runserver 0.0.0.0:8000`
5. Exposes your app on port 8000

---

## ✅ As a DevOps beginner, the order to write a Dockerfile is:

1. Choose a **base image** (`FROM`)
2. Set your **working directory** (`WORKDIR`)
3. **Copy dependencies** (`COPY requirements.txt ...`)
4. **Install dependencies** (`RUN pip install ...`)
5. **Copy your app code**
6. **Expose ports** (`EXPOSE`)
7. **Set startup command** (`ENTRYPOINT` + `CMD`)


```

Docker does **two things automatically**:

1. **If `/app` doesn’t exist**, it **creates** that directory inside the container.
2. Then it **switches** to that directory — just like running `cd /app` in a terminal.

So from that point onward, every command (`COPY`, `RUN`, etc.) runs inside `/app` unless you change it again with another `WORKDIR`.

### Example:

```dockerfile
FROM ubuntu

WORKDIR /app
RUN echo "This is inside /app" > test.txt

WORKDIR /code
RUN echo "Now inside /code" > file.txt
```

Inside the container, it will have:

```
/app/test.txt
/code/file.txt
```

That’s because each `WORKDIR` creates the folder if missing and changes into it.

So yes — your understanding is 100% correct:
**`WORKDIR` both creates and switches to that directory.**
