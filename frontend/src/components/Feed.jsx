import React, { useState, useEffect } from 'react';
import { useTheme } from '../contexts/ThemeContext';

const Feed = ({ user }) => {
  const { theme } = useTheme();
  const [posts, setPosts] = useState([]);
  const [newPost, setNewPost] = useState('');
  const [isPosting, setIsPosting] = useState(false);

  // Mock data for demonstration
  useEffect(() => {
    // Simulate loading posts
    const mockPosts = [
      {
        id: 1,
        username: 'john_doe',
        gender: 'male',
        content: 'Just joined The Gruv! Loving the blue theme 💙',
        timestamp: new Date(Date.now() - 3600000), // 1 hour ago
        likes: 12
      },
      {
        id: 2,
        username: 'jane_smith',
        gender: 'female',
        content: 'The pink theme is absolutely beautiful! Perfect for my aesthetic 💕',
        timestamp: new Date(Date.now() - 7200000), // 2 hours ago
        likes: 8
      },
      {
        id: 3,
        username: 'alex_purple',
        gender: 'other',
        content: 'Purple theme crew checking in! 💜 The color customization is amazing.',
        timestamp: new Date(Date.now() - 10800000), // 3 hours ago
        likes: 15
      },
      {
        id: 4,
        username: 'demo_user',
        gender: 'default',
        content: 'Welcome to The Gruv community! Share your thoughts and connect with others.',
        timestamp: new Date(Date.now() - 14400000), // 4 hours ago
        likes: 23
      }
    ];
    setPosts(mockPosts);
  }, []);

  const getThemeColorForGender = (gender) => {
    switch (gender) {
      case 'male': return '#2563eb';
      case 'female': return '#ec4899';
      case 'other': return '#7c3aed';
      default: return '#374151';
    }
  };

  const formatTimestamp = (timestamp) => {
    const now = new Date();
    const diff = now - timestamp;
    const minutes = Math.floor(diff / 60000);
    const hours = Math.floor(diff / 3600000);
    const days = Math.floor(diff / 86400000);

    if (minutes < 60) return `${minutes}m ago`;
    if (hours < 24) return `${hours}h ago`;
    return `${days}d ago`;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!newPost.trim()) return;

    setIsPosting(true);
    
    // Simulate API call
    setTimeout(() => {
      const post = {
        id: posts.length + 1,
        username: user?.user?.username || 'You',
        gender: user?.user?.gender || 'default',
        content: newPost,
        timestamp: new Date(),
        likes: 0
      };
      
      setPosts([post, ...posts]);
      setNewPost('');
      setIsPosting(false);
    }, 1000);
  };

  const handleLike = (postId) => {
    setPosts(posts.map(post => 
      post.id === postId 
        ? { ...post, likes: post.likes + 1 }
        : post
    ));
  };

  return (
    <div className="container" style={{ paddingTop: '2rem', paddingBottom: '2rem' }}>
      <div style={{ maxWidth: '600px', margin: '0 auto' }}>
        {/* Header */}
        <div style={{ marginBottom: '2rem', textAlign: 'center' }}>
          <h1 style={{ 
            fontSize: 'var(--font-size-3xl)', 
            fontWeight: 'bold', 
            color: 'var(--color-primary)',
            marginBottom: '0.5rem'
          }}>
            The Gruv Feed
          </h1>
          <p style={{ color: 'var(--color-textSecondary)' }}>
            Connect with the community and share your thoughts
          </p>
        </div>

        {/* New Post Form */}
        <div className="card" style={{ marginBottom: '2rem' }}>
          <form onSubmit={handleSubmit}>
            <div className="form-group">
              <textarea
                className="form-input"
                value={newPost}
                onChange={(e) => setNewPost(e.target.value)}
                placeholder="What's on your mind?"
                rows="3"
                style={{ 
                  resize: 'vertical', 
                  minHeight: '80px',
                  backgroundColor: 'var(--color-background)'
                }}
              />
            </div>
            <div style={{ 
              display: 'flex', 
              justifyContent: 'space-between', 
              alignItems: 'center'
            }}>
              <div style={{ 
                display: 'flex', 
                alignItems: 'center', 
                gap: '0.5rem',
                color: 'var(--color-textSecondary)',
                fontSize: 'var(--font-size-sm)'
              }}>
                <div style={{
                  width: '16px',
                  height: '16px',
                  borderRadius: '50%',
                  backgroundColor: 'var(--color-primary)'
                }} />
                Posting as {user?.user?.username}
              </div>
              <button
                type="submit"
                className="btn"
                disabled={!newPost.trim() || isPosting}
                style={{ 
                  opacity: (!newPost.trim() || isPosting) ? 0.7 : 1,
                  cursor: (!newPost.trim() || isPosting) ? 'not-allowed' : 'pointer'
                }}
              >
                {isPosting ? 'Posting...' : 'Post'}
              </button>
            </div>
          </form>
        </div>

        {/* Posts Feed */}
        <div style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
          {posts.length === 0 ? (
            <div className="card" style={{ textAlign: 'center', padding: '3rem 1.5rem' }}>
              <p style={{ 
                color: 'var(--color-textSecondary)', 
                fontSize: 'var(--font-size-lg)' 
              }}>
                No posts yet. Be the first to share something!
              </p>
            </div>
          ) : (
            posts.map((post) => (
              <div key={post.id} className="card">
                {/* Post Header */}
                <div style={{ 
                  display: 'flex', 
                  alignItems: 'center', 
                  gap: '1rem', 
                  marginBottom: '1rem'
                }}>
                  <div style={{
                    width: '40px',
                    height: '40px',
                    borderRadius: '50%',
                    backgroundColor: getThemeColorForGender(post.gender),
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    color: 'white',
                    fontWeight: 'bold',
                    fontSize: 'var(--font-size-sm)'
                  }}>
                    {post.username[0].toUpperCase()}
                  </div>
                  <div>
                    <h3 style={{ 
                      fontWeight: '600', 
                      fontSize: 'var(--font-size-base)',
                      color: 'var(--color-text)',
                      marginBottom: '0.25rem'
                    }}>
                      {post.username}
                    </h3>
                    <p style={{ 
                      color: 'var(--color-textSecondary)', 
                      fontSize: 'var(--font-size-sm)'
                    }}>
                      {formatTimestamp(post.timestamp)}
                    </p>
                  </div>
                </div>

                {/* Post Content */}
                <div style={{ marginBottom: '1rem' }}>
                  <p style={{ 
                    color: 'var(--color-text)', 
                    lineHeight: '1.6',
                    fontSize: 'var(--font-size-base)'
                  }}>
                    {post.content}
                  </p>
                </div>

                {/* Post Actions */}
                <div style={{ 
                  display: 'flex', 
                  alignItems: 'center', 
                  gap: '1rem',
                  paddingTop: '0.75rem',
                  borderTop: '1px solid var(--color-border)'
                }}>
                  <button
                    onClick={() => handleLike(post.id)}
                    style={{
                      display: 'flex',
                      alignItems: 'center',
                      gap: '0.5rem',
                      background: 'none',
                      border: 'none',
                      color: 'var(--color-textSecondary)',
                      fontSize: 'var(--font-size-sm)',
                      cursor: 'pointer',
                      padding: '0.25rem 0.5rem',
                      borderRadius: 'var(--radius-base)',
                      transition: 'all 0.2s ease'
                    }}
                    onMouseOver={(e) => {
                      e.target.style.backgroundColor = 'var(--color-background)';
                      e.target.style.color = 'var(--color-primary)';
                    }}
                    onMouseOut={(e) => {
                      e.target.style.backgroundColor = 'transparent';
                      e.target.style.color = 'var(--color-textSecondary)';
                    }}
                  >
                    ♥ {post.likes}
                  </button>
                </div>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
};

export default Feed;