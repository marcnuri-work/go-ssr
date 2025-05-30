// https://thoughtbot.github.io/superglue/recipes/ssr/
export {TextEncoder, TextDecoder} from 'text-encoding'

const messageChannel = function () {
  this.port1 = {
    postMessage: function (message) {},
  };

  this.port2 = {
    addEventListener: function (event, handler) {
      this._eventHandler = handler;
    },
    removeEventListener: function (event) {
      this._eventHandler = null;
    },
    simulateMessage: function (data) {
      if (this._eventHandler) {
        this._eventHandler({ data });
      }
    },
  };
};

export const MessageChannel = typeof document != "undefined" ? window.MessageChannel : messageChannel;
