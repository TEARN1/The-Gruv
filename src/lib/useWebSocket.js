import { useEffect, useRef, useState } from 'react';

export default function useWebSocket() {
  const wsRef = useRef(null);
  const listenersRef = useRef(new Set());
  const [lastMessage, setLastMessage] = useState(null);
  const reconnectRef = useRef(null);

  useEffect(() => {
    const url = import.meta.env.VITE_WS_URL || 'ws://localhost:8080';

    function connect() {
      const ws = new WebSocket(url);
      wsRef.current = ws;

      ws.onopen = () => {
        console.debug('[WS] connected', url);
      };

      ws.onmessage = (ev) => {
        try {
          const data = JSON.parse(ev.data);
          setLastMessage(data);
          for (const fn of listenersRef.current) fn(data);
        } catch (e) {
          console.warn('[WS] invalid json', ev.data);
        }
      };

      ws.onclose = () => {
        console.debug('[WS] closed, reconnecting in 2s');
        // try reconnect
        reconnectRef.current = setTimeout(connect, 2000);
      };

      ws.onerror = (e) => {
        console.debug('[WS] error', e);
        ws.close();
      };
    }

    connect();
    return () => {
      if (reconnectRef.current) clearTimeout(reconnectRef.current);
      if (wsRef.current) wsRef.current.close();
      listenersRef.current.clear();
    };
  }, []);

  function sendMessage(obj) {
    if (wsRef.current && wsRef.current.readyState === WebSocket.OPEN) {
      wsRef.current.send(JSON.stringify(obj));
      return true;
    }
    return false;
  }

  function subscribe(fn) {
    listenersRef.current.add(fn);
    return () => listenersRef.current.delete(fn);
  }

  return { sendMessage, subscribe, lastMessage };
}