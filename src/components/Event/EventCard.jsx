import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { postComment } from '../../lib/api';

export default function EventCard({ event }) {
  const nav = useNavigate();
  const [comment, setComment] = useState('');
  const [posting, setPosting] = useState(false);
  const submit = async (e) => {
    e.preventDefault();
    if (!comment.trim()) return;
    try {
      setPosting(true);
      await postComment(event.id, comment);
      setComment('');
      // UI will update via WebSocket broadcast
    } catch (err) {
      console.error(err);
    } finally { setPosting(false); }
  };

  return (
    <div className="card" onClick={() => nav(`/event/${event.id}`)}>
      <div className="card-header">
        <strong>{event.author?.name || 'Unknown'}</strong>
        <span>{new Date(event.postedAt || Date.now()).toLocaleString()}</span>
      </div>
      <h3 className="title">{event.title}</h3>
      <p className="muted">{event.description}</p>
      <div className="card-footer">
        <button onClick={(e) => { e.stopPropagation(); alert('RSVP (placeholder)'); }}>RSVP</button>
        <button onClick={(e) => { e.stopPropagation(); alert('Amplify (placeholder)'); }}>Amplify</button>
      </div>
      <div className="comments" onClick={(e) => e.stopPropagation()}> 
        {event.comments?.slice(-3).map(c => <div key={c.id}><b>{c.userName}</b>: {c.text}</div>)}
        <form onSubmit={submit} className="comment-form">
          <input value={comment} onChange={(e) => setComment(e.target.value)} placeholder="Add a comment..." />
          <button type="submit" disabled={posting}>{posting ? '...' : 'Send'}</button>
        </form>
      </div>
    </div>
  );
}