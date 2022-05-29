import * as Y from 'yjs'
import {syncronize} from 'y-pojo'
import { fromUint8Array, toUint8Array } from 'js-base64'

var doc
var root
const key = "d"

export const initialize = () => {
    doc = new Y.Doc()
    root = doc.getMap("r")
    if (documentText && documentText.length > 0) {
        root.set("_t", new Y.Text())
        root.get("_t").insert(0, documentText)
    } else if (documentObject !== undefined) {
        syncronize(root, JSON.parse(documentObject))
    }

    return "initialized"
}

export const applyUpdate = () => {
    let data = toUint8Array(encodedUpdate)
    Y.applyUpdateV2(doc, data)

    return "hello"
}

export const encodeStateAsUpdate = () => {
    let stateVector = undefined
    if (encodedStateVector && encodedStateVector.length > 0) {
        stateVector = toUint8Array(encodedStateVector)
    }
    let arr = Y.encodeStateAsUpdateV2(doc, stateVector)
    return fromUint8Array(arr)
}

export const stateVector = () => {
    return fromUint8Array(Y.encodeStateVector(doc))
}

export const toString = () => {
    return root.get("_t").toString()
}

export const toJSON = () => {
    return JSON.stringify(root.toJSON())
}

// Server doesn't actually modify the document, these are for testing

export const insert = () => {
    root.get("d").insert(insertPosition, insertText)
}