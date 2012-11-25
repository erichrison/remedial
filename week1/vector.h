// A dynamically resized array.
#pragma once

#include <cstddef>  // for size_t
#include <iterator>  // for reverse_iterator

// "Remedial Template Library"
namespace rtl {

// Other than the typedef's at the top, I believe that this is pretty much
// the STL's vector class. It seemed like a good starting point.
template <typename T>
class vector {
 public:
  typedef size_t size_type;
  typedef T* iterator;
  typedef const T* const_iterator;
  typedef std::reverse_iterator<iterator> reverse_iterator;
  typedef std::reverse_iterator<const_iterator> const_reverse_iterator;

  explicit vector(size_type n = 0, const T& x = T());
  vector(const vector<T>& v);
  template <typename I> vector(I first, I last);

  iterator begin();
  const_iterator begin() const;

  iterator end();
  const_iterator end() const;

  reverse_iterator rbegin();
  const_reverse_iterator rbegin() const;

  reverse_iterator rend();
  const_reverse_iterator rend() const;

  size_type size() const;
  size_type max_size() const;
  void resize(size_type size, const T& copy = T());
  size_type capacity() const;
  bool empty() const;
  void reserve(size_type n);

  T& operator[](size_type n);
  const T& operator[](size_type n) const;
  T& at(size_type n);
  const T& at(size_type n) const;
  T& front();
  const T& front() const;
  T& back();
  const T& back() const;

  template <typename I> void assign(I first, I last);
  void assign(size_type n, const T& x);
  void push_back(const T& x);
  void pop_back();
  iterator insert(iterator p, const T& x);
  iterator insert(iterator p, size_type n, const T& x);
  template <typename I> iterator insert(iterator p, I first, I last);
  iterator erase(iterator p);
  iterator erase(iterator first, iterator last);
  void swap(vector<T>& v);
  void clear();
};

// Add some code here, probably

}  // namespace rtl
