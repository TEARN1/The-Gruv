import React, { useState } from 'react';
import { useTheme } from '../contexts/ThemeContext';
import { authAPI } from '../utils/api';

const Profile = ({ user, onUpdate }) => {
  const { theme, updateTheme } = useTheme();
  const [isEditing, setIsEditing] = useState(false);
  const [formData, setFormData] = useState({
    gender: user?.user?.gender || '',
    email: user?.user?.email || '',
    bio: user?.user?.bio || '',
    avatarUrl: user?.user?.avatarUrl || ''
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value
    });
    
    // Preview theme changes when gender is updated
    if (name === 'gender' && value) {
      updateTheme(value);
    }
    
    // Clear messages
    if (error) setError('');
    if (success) setSuccess('');
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');
    setSuccess('');

    try {
      const response = await authAPI.updateProfile(user.userId, formData);
      onUpdate(response);
      setSuccess('Profile updated successfully!');
      setIsEditing(false);
    } catch (err) {
      setError(err.error || 'Failed to update profile. Please try again.');
      // Reset theme if update failed
      updateTheme(user?.user?.gender || 'default');
    } finally {
      setIsLoading(false);
    }
  };

  const handleCancel = () => {
    // Reset form to original values
    setFormData({
      gender: user?.user?.gender || '',
      email: user?.user?.email || '',
      bio: user?.user?.bio || '',
      avatarUrl: user?.user?.avatarUrl || ''
    });
    // Reset theme
    updateTheme(user?.user?.gender || 'default');
    setIsEditing(false);
    setError('');
    setSuccess('');
  };

  const currentGender = user?.user?.gender || 'default';
  const genderLabels = {
    male: 'Male (Blue Theme)',
    female: 'Female (Pink Theme)', 
    other: 'Other (Purple Theme)',
    default: 'No Preference (Default Theme)'
  };

  return (
    <div className="container" style={{ paddingTop: '2rem', paddingBottom: '2rem' }}>
      <div style={{ maxWidth: '800px', margin: '0 auto' }}>
        {/* Header */}
        <div style={{ marginBottom: '2rem', textAlign: 'center' }}>
          <h1 style={{ 
            fontSize: 'var(--font-size-3xl)', 
            fontWeight: 'bold', 
            color: 'var(--color-primary)',
            marginBottom: '0.5rem'
          }}>
            Your Profile
          </h1>
          <p style={{ color: 'var(--color-textSecondary)' }}>
            Manage your account settings and preferences
          </p>
        </div>

        {/* Profile Card */}
        <div className="card">
          {/* Avatar Section */}
          <div style={{ 
            display: 'flex', 
            alignItems: 'center', 
            gap: '1.5rem', 
            marginBottom: '2rem',
            paddingBottom: '1.5rem',
            borderBottom: '1px solid var(--color-border)'
          }}>
            <div style={{
              width: '80px',
              height: '80px',
              borderRadius: '50%',
              backgroundColor: 'var(--color-primary)',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              color: 'white',
              fontSize: 'var(--font-size-2xl)',
              fontWeight: 'bold',
              backgroundImage: formData.avatarUrl ? `url(${formData.avatarUrl})` : 'none',
              backgroundSize: 'cover',
              backgroundPosition: 'center'
            }}>
              {!formData.avatarUrl && (user?.user?.username?.[0]?.toUpperCase() || 'U')}
            </div>
            <div>
              <h2 style={{ 
                fontSize: 'var(--font-size-xl)', 
                fontWeight: 'bold', 
                marginBottom: '0.25rem',
                color: 'var(--color-text)'
              }}>
                {user?.user?.username || 'Unknown User'}
              </h2>
              <p style={{ 
                color: 'var(--color-textSecondary)',
                fontSize: 'var(--font-size-sm)'
              }}>
                Current Theme: {genderLabels[currentGender]}
              </p>
            </div>
          </div>

          {/* Profile Form */}
          <form onSubmit={handleSubmit}>
            <div style={{ 
              display: 'grid', 
              gap: '1.5rem',
              gridTemplateColumns: window.innerWidth > 768 ? '1fr 1fr' : '1fr'
            }}>
              <div className="form-group">
                <label htmlFor="gender" className="form-label">
                  Gender Preference (Theme)
                </label>
                <select
                  id="gender"
                  name="gender"
                  className="form-select"
                  value={formData.gender}
                  onChange={handleChange}
                  disabled={!isEditing}
                >
                  <option value="">Default Theme</option>
                  <option value="male">Male (Blue Theme)</option>
                  <option value="female">Female (Pink Theme)</option>
                  <option value="other">Other (Purple Theme)</option>
                </select>
                {isEditing && (
                  <p style={{ 
                    fontSize: 'var(--font-size-sm)', 
                    color: 'var(--color-textSecondary)',
                    marginTop: '0.25rem'
                  }}>
                    Changes preview immediately
                  </p>
                )}
              </div>

              <div className="form-group">
                <label htmlFor="email" className="form-label">
                  Email
                </label>
                <input
                  type="email"
                  id="email"
                  name="email"
                  className="form-input"
                  value={formData.email}
                  onChange={handleChange}
                  disabled={!isEditing}
                  placeholder="your@email.com"
                />
              </div>
            </div>

            <div className="form-group">
              <label htmlFor="bio" className="form-label">
                Bio
              </label>
              <textarea
                id="bio"
                name="bio"
                className="form-input"
                value={formData.bio}
                onChange={handleChange}
                disabled={!isEditing}
                placeholder="Tell us about yourself..."
                rows="4"
                style={{ resize: 'vertical', minHeight: '100px' }}
              />
            </div>

            <div className="form-group">
              <label htmlFor="avatarUrl" className="form-label">
                Avatar URL
              </label>
              <input
                type="url"
                id="avatarUrl"
                name="avatarUrl"
                className="form-input"
                value={formData.avatarUrl}
                onChange={handleChange}
                disabled={!isEditing}
                placeholder="https://example.com/your-avatar.jpg"
              />
            </div>

            {/* Messages */}
            {error && (
              <div className="error-message" style={{ marginBottom: '1rem' }}>
                {error}
              </div>
            )}

            {success && (
              <div className="success-message" style={{ marginBottom: '1rem' }}>
                {success}
              </div>
            )}

            {/* Action Buttons */}
            <div style={{ 
              display: 'flex', 
              gap: '1rem', 
              justifyContent: 'flex-end',
              paddingTop: '1rem',
              borderTop: '1px solid var(--color-border)'
            }}>
              {!isEditing ? (
                <button
                  type="button"
                  className="btn"
                  onClick={() => setIsEditing(true)}
                >
                  Edit Profile
                </button>
              ) : (
                <>
                  <button
                    type="button"
                    className="btn btn-secondary"
                    onClick={handleCancel}
                    disabled={isLoading}
                  >
                    Cancel
                  </button>
                  <button
                    type="submit"
                    className="btn"
                    disabled={isLoading}
                    style={{ 
                      opacity: isLoading ? 0.7 : 1,
                      cursor: isLoading ? 'not-allowed' : 'pointer'
                    }}
                  >
                    {isLoading ? 'Saving...' : 'Save Changes'}
                  </button>
                </>
              )}
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Profile;