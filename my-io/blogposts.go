package my_io

import "io/fs"

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	entries, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range entries {
		post, err := getPost(fileSystem, f)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, entry fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(entry.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}
