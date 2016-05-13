/**
 * Created by Administrator on 2016/5/9.
 */
(function() {

    var format = require("string-format");

    var logger = {
        Debug : function () {
            var msg = format.apply(null, arguments);
            console.log(new Date().toLocaleString() + " Debug: [Javascript] " + msg);
        },
        Info : function () {
            var msg = format.apply(null, arguments);
            console.log(new Date().toLocaleString() + " Info: [Javascript] " + msg);
        },
        Warning : function () {
            var msg = format.apply(null, arguments);
            console.log(new Date().toLocaleString() + " Warning: [Javascript] " + msg);
        },
        Error : function () {
            var msg = format.apply(null, arguments);
            console.log(new Date().toLocaleString() + " Error: [Javascript] " + msg);
        },
        Fatal : function () {
            var msg = format.apply(null, arguments);
            console.log(new Date().toLocaleString() + " Fatal: [Javascript] " + msg);
        },
    }

    if (typeof module !== 'undefined') {
        module.exports = logger;
    } else if (typeof define === 'function' && define.amd) {
        define(logger);
    } else {
        window.logger = logger;
    }

}).call(this);
