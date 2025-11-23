import React, { useEffect, useState } from 'react';
import { fetchForYou } from '../../lib/api';
import EventCard from '../Event/EventCard';
import useWebSocket from '../../lib/useWebSocket';

export default function Feed() {
  const [events, setEvents] = useState([]);
  const [loading, setLoading] = useState(true);
  const [err, setErr] = useState(null);
  const { subscribe } = useWebSocket();

  useEffect(() => {
    let mounted = true;
    fetchForYou()
      .then((data) => { if (mounted) setEvents(data); })
      .catch((e) => setErr(e.message))
      .finally(() => setLoading(false));
    return () => { mounted = false; };
  }, []);

  useEffect(() => {
    // listen for server broadcasts (NEW_COMMENT etc.)
    const unsub = subscribe((msg) => {
      if (!msg || !msg.type) return;
      if (msg.type === 'NEW_COMMENT') {
        const { happeningId, comment } = msg;
        setEvents((prev) => prev.map((ev) => ev.id === happeningId ? { ...ev, comments: [...(ev.comments || []), comment] } : ev));
      }
      // handle other message types later (SIGNAL_UPDATE, GRUVDROP_UPDATE)
    });
    return unsub;
  }, [subscribe]);

  if (loading) return <div className="centered">Loading The Pulse...</div>;
  if (err) return <div className="centered error">Error: {err}</div>;

  return (
    <div className="container">
      <h2>The Pulse</h2>
      <div className="list">
        {events.map((ev) => <EventCard key={ev.id} event={ev} />)}
      </div>
    </div>
  );
}