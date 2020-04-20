FROM node:latest

# creating work directory.
WORKDIR /usr/src/app

# copying file with required dependencies in work directory.
COPY package*.json ./

# installing all required dependencies.
RUN npm install

# copying all files to work directory.
COPY . .

# exposing port.
EXPOSE 5000

# adding new host.
ENV HOST=0.0.0.0

# running the app.
CMD ["npm", "run", "dev"]