/**
 * Created by Administrator on 2016/5/9.
 */

var logger = {
    Debug : function (msg) {
        var args = arguments, re = new RegExp("%([1-" + args.length + "])", "g");
        console.log(new Date().toLocaleString() + " Debug: [Javascript] " + String(msg).replace(re,
                function($1, $2) {
                    return args[$2];
                }));
    },
    Info : function (args) {
        var i = 1;
        var msg = args[0].replace(/%s/g, function(){
            return args[i++];
        });
        console.log(new Date().toLocaleString() + " Info: [Javascript] " + msg);
    },
    Warning : function (args) {
        var i = 1;
        var msg = args[0].replace(/%s/g, function(){
            return args[i++];
        });
        console.log(new Date().toLocaleString() + " Warning: [Javascript] " + msg);
    },
    Error : function (args) {
        var i = 1;
        var msg = args[0].replace(/%s/g, function(){
            return args[i++];
        });
        console.log(new Date().toLocaleString() + " Error: [Javascript] " + msg);
    },
    Fatal : function (args) {
        var i = 1;
        var msg = args[0].replace(/%s/g, function(){
            return args[i++];
        });
        console.log(new Date().toLocaleString() + " Fatal: [Javascript] " + msg);
    },
}