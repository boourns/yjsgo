import * as Y from 'yjs'
import {syncronize} from 'y-pojo'
import { fromUint8Array, toUint8Array } from 'js-base64'

var doc
var root

export const initialize = () => {
    doc = new Y.Doc()
    if (complex) {
        root = doc.getMap("r")
    } else {
        root = doc.getText("t")
    }

    if (!complex && documentText && documentText.length > 0) {
        root.insert(0, documentText)
    } else if (complex && documentObject !== undefined) {
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
    if (complex) {
        return JSON.stringify(root.toJSON())
    } else {
        return root.toString()
    }
}

// Server doesn't actually modify the document, these are for testing

export const insert = () => {
    root.insert(insertPosition, insertText)
}