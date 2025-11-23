import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { fetchForYou } from '../../lib/api';

export default function EventDetail() {
  const { id } = useParams();
  const [event, setEvent] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let active = true;
    // Simple fetch: reuse forYou and pick item by id for now
    fetchForYou()
      .then(list => { if (active) setEvent(list.find(e => e.id === id)); })
      .catch(console.error)
      .finally(() => active && setLoading(false));
    return () => { active = false; };
  }, [id]);

  if (loading) return <div className="centered">Loading event...</div>;
  if (!event) return <div className="centered">Event not found</div>;

  return (
    <div className="container">
      <h2>{event.title}</h2>
      <p className="muted">by {event.author?.name}</p>
      <div className="hero">
        <img src={event.images?.[0] || '/placeholder.png'} alt={event.title} />
      </div>
      <section>
        <h3>The Blueprint</h3>
        {event.blueprint?.map(ch => <div key={ch.chapter}><h4>{ch.title}</h4><p>{ch.content}</p></div>)}
      </section>
      <section>
        <h3>Concierge</h3>
        <div className="grid">
          <a className="concierge" href={`https://m.uber.com/ul/?action=setPickup&pickup=my_location&dropoff[formatted_address]=${encodeURIComponent(event.address || '')}`} target="_blank">Catch a Ride</a>
          <a className="concierge" href={`https://www.booking.com/searchresults.html?ss=${encodeURIComponent(event.suburb || '')}`} target="_blank">Stay</a>
          <button disabled className="concierge coming-soon">Find the Anthem (Coming Soon)</button>
        </div>
      </section>
    </div>
  );
}