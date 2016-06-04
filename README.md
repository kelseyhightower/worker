# worker

This is not an official Google project.

worker pulls messages from a memq queue for testing distributed infrastructure deployments. This is a prototype and not ment for use outside of my workshops and demos.

## Usage

```
worker -queue=work --memq=http://127.0.0.1:80
```

```
2016/06/04 12:09:22 Starting worker...
2016/06/04 12:09:22 Work queue set to work
2016/06/04 12:09:22 Processing message id: 7c5d34ce-da61-c161-86bf-e74b54d6ba02
2016/06/04 12:09:23 Processing message id: 5b511a48-5d00-9141-605d-8a60eaf07b1d
2016/06/04 12:09:24 Processing message id: 2cf97605-03e6-e467-0d36-41ebb7f3c4b1
2016/06/04 12:09:25 Work queue is empty. Shutting down...
```
