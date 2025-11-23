import React, { useEffect, useState } from 'react';
import { fetchForYou } from '../../lib/api';
import EventCard from '../Event/EventCard';

export default function Feed() {
  const [events, setEvents] = useState([]);
  const [loading, setLoading] = useState(true);
  const [err, setErr] = useState(null);

  useEffect(() => {
    let mounted = true;
    fetchForYou()
      .then((data) => { if (mounted) setEvents(data); })
      .catch((e) => setErr(e.message))
      .finally(() => setLoading(false));
    return () => { mounted = false; };
  }, []);

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