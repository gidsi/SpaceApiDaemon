import React from 'react';
import PropTypes from 'prop-types';
import injectSheet from 'react-jss';
import classnames from 'classnames';

const style = {
  container: {
    position: 'relative',
    width: '100%',
  },
  label: {
    position: 'absolute',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    height: '100%',
    width: '100%',
    backgroundColor: 'transparent',
    border: '1px #eee solid',
    top: '0px',
    borderRadius: '5px 5px 0 0',
    left: '0px',
    '&:hover': {
      backgroundColor: '#eee',
      borderColor: '#eee',
    },
  },
  checked: {
    backgroundColor: '#ccc',
    color: '#fff',
    '&:hover': {
      backgroundColor: '#ccc',
      color: '#fff',
    },
  },
  input: {
    visibility: 'hidden',
  },
};

const RadioButton = props => (
  <div className={props.classes.container}>
    <input
      type="radio"
      id={props.id}
      name={props.group}
      onClick={props.onChange}
      onChange={props.onChange}
      className={props.classes.input}
      value={props.value}
      checked={props.checked}
    />
    <label
      htmlFor={props.id}
      className={classnames(props.classes.label, { [props.classes.checked]: props.checked })}
    >
      {props.display}
    </label>
  </div>
);

RadioButton.propTypes = {
  id: PropTypes.string.isRequired,
  value: PropTypes.string.isRequired,
  checked: PropTypes.bool.isRequired,
  onChange: PropTypes.func,
  display: PropTypes.string,
  group: PropTypes.string,
  classes: PropTypes.object.isRequired, // eslint-disable-line react/forbid-prop-types
};

RadioButton.defaultProps = {
  display: '',
  group: 'radio',
  onChange: () => {},
};

export default injectSheet(style)(RadioButton);
