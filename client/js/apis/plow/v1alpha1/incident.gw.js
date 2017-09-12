// Code generated by protoc-gen-grpc-js-client
// source: incident.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function incidentsList(p, conf) {
    path = '/plow/v1alpha1/incidents/json'
    return xhr(path, 'GET', conf, p);
}

function incidentsDescribe(p, conf) {
    path = '/plow/v1alpha1/incidents/' + p['phid'] + '/json'
    delete p['phid']
    return xhr(path, 'GET', conf, p);
}

function incidentsNotify(p, conf) {
    path = '/plow/v1alpha1/clusters/' + p['kubernetes_cluster'] + '/actions/notify-incident/json'
    delete p['kubernetes_cluster']
    return xhr(path, 'POST', conf, null, p);
}

function incidentsCreateEvent(p, conf) {
    path = '/plow/v1alpha1/incidents/' + p['phid'] + '/events/json'
    delete p['phid']
    return xhr(path, 'POST', conf, null, p);
}

var services = {
    incidents: {
        list: incidentsList,
        describe: incidentsDescribe,
        notify: incidentsNotify,
        createEvent: incidentsCreateEvent
    }
};

module.exports = {appscode: {plow: {v1alpha1: services}}};
