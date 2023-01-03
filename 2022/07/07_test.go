package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const input = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

/*
- / (dir)
  - a (dir)
    - e (dir)
      - i (file, size=584)
    - f (file, size=29116)
    - g (file, size=2557)
    - h.lst (file, size=62596)
  - b.txt (file, size=14848514)
  - c.dat (file, size=8504156)
  - d (dir)
    - j (file, size=4060174)
    - d.log (file, size=8033020)
    - d.ext (file, size=5626152)
    - k (file, size=7214296)
*/

var expected = Directory{
	Name:   "/",
	Parent: nil,
	Childs: []*Directory{
		{
			Name: "a",
			Childs: []*Directory{
				{
					Name: "e",
					Files: []*File{
						{
							Name: "i",
							Size: 584,
						},
					},
				},
			},
			Files: []*File{
				{
					Name: "f",
					Size: 29116,
				},
				{
					Name: "g",
					Size: 2557,
				},
				{
					Name: "h.lst",
					Size: 62596,
				},
			},
		},
		{
			Name: "d",
			Files: []*File{
				{
					Name: "j",
					Size: 4060174,
				},
				{
					Name: "d.log",
					Size: 8033020,
				},
				{
					Name: "d.ext",
					Size: 5626152,
				},
				{
					Name: "k",
					Size: 7214296,
				},
			},
		},
	},
	Files: []*File{
		{
			Name: "b.txt",
			Size: 14848514,
		},
		{
			Name: "c.dat",
			Size: 8504156,
		},
	},
}

func (d *Directory) cleanParent() {
	d.Parent = nil
	for _, c := range d.Childs {
		c.cleanParent()
	}
}

func TestParseTree(t *testing.T) {
	tree, err := parseTree([]byte(input))
	if err != nil {
		t.Fatalf("parseTree(); err=%v, want nil", err)
	}

	// clear out parent pointer for comparison
	tree.cleanParent()

	if diff := cmp.Diff(&expected, tree); diff != "" {
		t.Errorf("parseTree() mismatch (-want +got):\n%s", diff)
	}
	// The total size of directory e is 584 because it contains a single file i of size 584 and no other directories.
	checkDirSize(t, tree.Childs[0].Childs[0], 584)
	// The directory a has total size 94853 because it contains files f (size 29116), g (size 2557), and h.lst (size 62596), plus file i indirectly (a contains e which contains i).
	checkDirSize(t, tree.Childs[0], 94853)
	// Directory d has total size 24933642.
	checkDirSize(t, tree.Childs[1], 24933642)
	// As the outermost directory, / contains every file. Its total size is 48381165, the sum of the size of every file.
	checkDirSize(t, tree, 48381165)
}

func checkDirSize(t *testing.T, d *Directory, want int) {
	got := d.calculateSize()
	if got != want {
		t.Errorf("%s.calculateSize() got %d, want %d", d.Name, got, want)
	}
}
