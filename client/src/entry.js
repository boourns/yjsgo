import * as Y from 'yjs'
import { fromUint8Array, toUint8Array } from 'js-base64'

var doc
const key = "d"

export const initialize = () => {
    doc = new Y.Doc()
    if (documentText && documentText.length > 0) {
        doc.getText().insert(0, documentText)
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
    return doc.getText().toString()
}

// Server doesn't actually modify the document, these are for testing

export const insert = () => {
    doc.getText().insert(insertPosition, insertText)
}