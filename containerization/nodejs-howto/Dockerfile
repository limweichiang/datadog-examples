FROM node:14

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./
RUN npm install

# Install Datadog Node.js tracer. This is shown here explicitly as an example.
# Feel free to include this as a dependency in package.json if needed. 
RUN npm install dd-trace --save

# Bundle app source
COPY . .

EXPOSE 8080
CMD [ "node", "--require", "dd-trace/init", "server.js" ]

