// Code generated by protoc-gen-grpc-js-client
// source: cluster.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function clustersList(p, conf) {
    path = '/plow/v1alpha1/clusters/json'
    return xhr(path, 'GET', conf, p);
}

function clustersDescribe(p, conf) {
    path = '/plow/v1alpha1/clusters/' + p['uid'] + '/json'
    delete p['uid']
    return xhr(path, 'GET', conf, p);
}

function clustersCreate(p, conf) {
    path = '/plow/v1alpha1/clusters/json'
    return xhr(path, 'POST', conf, null, p);
}

function clustersUpdate(p, conf) {
    path = '/plow/v1alpha1/clusters/' + p['name'] + '/json'
    delete p['name']
    return xhr(path, 'PUT', conf, null, p);
}

function clustersReconfigure(p, conf) {
    path = '/plow/v1alpha1/clusters/' + p['name'] + '/actions/reconfigure/json'
    delete p['name']
    return xhr(path, 'PUT', conf, null, p);
}

function clustersApply(p, conf) {
    path = '/plow/v1alpha1/clusters/' + p['name'] + '/actions/apply/json'
    delete p['name']
    return xhr(path, 'PUT', conf, null, p);
}

function clustersDelete(p, conf) {
    path = '/plow/v1alpha1/clusters/' + p['name'] + '/json'
    delete p['name']
    return xhr(path, 'DELETE', conf, p);
}

function clustersClientConfig(p, conf) {
    path = '/plow/v1alpha1/clusters/' + p['name'] + '/client-config/json'
    delete p['name']
    return xhr(path, 'GET', conf, p);
}

function clustersMetadata(p, conf) {
    path = '/plow/v1alpha1/clusters/' + p['uid'] + '/metadata/json'
    delete p['uid']
    return xhr(path, 'GET', conf, p);
}

var services = {
    clusters: {
        list: clustersList,
        describe: clustersDescribe,
        create: clustersCreate,
        update: clustersUpdate,
        reconfigure: clustersReconfigure,
        apply: clustersApply,
        delete: clustersDelete,
        clientConfig: clustersClientConfig,
        metadata: clustersMetadata
    }
};

module.exports = {appscode: {plow: {v1alpha1: services}}};
