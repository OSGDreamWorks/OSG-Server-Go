/**
 * Created by Administrator on 2016/5/13.
 */
(function() {
    var GameServer = osg.Class.extend({
        ctor:function (id) {
            logger.Debug("ctor");
        },
        addPlayer : function (cId, player) {
            logger.Debug("addPlayer : " + cId);
        },
        delPlayer : function (cId) {
            logger.Debug("delPlayer : " + cId);
        },
    });

    if (typeof module !== 'undefined') {
        module.exports = GameServer;
    } else if (typeof define === 'function' && define.amd) {
        define(GameServer);
    } else {
        window.GameServer = GameServer;
    }

}).call(this);