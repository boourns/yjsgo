# yjsgo
Run yjs from Go via v8

This is a proof-of-concept go module that can apply changes to y.js documents and generate updates.  It is not a port of y.js - it is running y.js inside a v8go context, inside the go application.

Note that this is designed for server architectures that run persistent goroutines for every active editing session, and there is no locking to protect access of the yjs.Document instance.  Or one could use this as an offline document "flattener" that applies a queue of YJS updates from clients to the server's copy of the document.

I have a pretty limited use-case for this library at this time, however would be interested in contributions that resolve open issues or ideas to improve it.
