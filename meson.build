project('tapocalc', 'cpp')

opt_bundle = get_option('bundle')
if opt_bundle
	rpath = '.'
else
	rpath = ''
endif

lib_src = ['src/libthapo-calc/libthapo-calc.cpp', 'src/libthapo-calc/grammar.cpp']
lib = shared_library('thapo-calc', sources : lib_src, version : '0.2', soversion : '0.2', install : true)
incdir = include_directories('src/libthapo-calc')
exec_src = ['src/thapo-calc-cli/thapo-calc-cli.cpp']

executable('thapo-calc-cli', sources : exec_src,  include_directories : incdir, link_with : lib, install : true, install_rpath : rpath)
install_data('README.md', install_dir : join_paths(get_option('datadir'), 'thapo-calc'))
