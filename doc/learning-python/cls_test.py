import unittest

'''
测试文件名可以以 test 开始或结尾
测试方法必须以 test_ 开头
'''


class TestStringMethods(unittest.TestCase):
    '''所有测试类继承自 unittest.TestCase'''

    longMessage = True
    '''
    longMessage: 控制自定义失败消息展示形式
        默认值为 True,
            如果传入 msg 参数, 自定义消息将附加到标准错误消息之后
        如果为 False,
            如果传入 msg 参数, 自定义消息将替换标准错误信息
            如果未传入 msg 参数, 使用标准错误信息
    '''

    def setUp(self):
        '''每个方法执行前都会执行'''
        print('test setUp...')
        # 输出测试方法的 id
        print(f'self.id() {self.id()}')
        # 输出测试方法的简要描述信息
        print(f'self.shortDescription() {self.shortDescription()}')

    def tearDown(self):
        '''每个方法执行后都会执行'''
        print('test tearDown...')

    # 条件为 真 时才跳过
    @unittest.skipIf(1 == 1, 'skip test_upper if the condition is true')
    def test_upper(self):
        self.assertEqual('foo'.upper(), 'FOO')
        self.assertNotEqual('foo'.upper(), 'foo')

    # 条件为 假 时才跳过
    @unittest.skipUnless(1 == 2, 'skip test_split unless the condition is true')
    def test_split(self):
        s = 'hello world'
        s_list = s.split(' ')
        self.assertIn('hello', s_list)
        self.assertNotIn('world!', s_list)

    # 跳过这个测试方法
    @unittest.skip('skip test_isupper')
    def test_isupper(self):
        self.assertTrue('FOO'.isupper())
        self.assertFalse('Foo'.isupper())

    # 比较两个值，第一个参数是否小于第二个参数
    def test_less_equal(self):
        self.assertLess(1, 2)
        self.assertLessEqual(1, 1)

    # 比较两个值，第一个参数是否大于第二个参数
    def test_greater_equal(self):
        self.assertGreater(2, 1)
        self.assertGreaterEqual(2, 2)

    # 比较两个多行字符串是否相等
    def test_multiline_str_equal(self):
        '''比较两个多行字符串是否相等'''
        self.assertMultiLineEqual('''
        hello world
        hello python
        ''', '''
        hello world
        ''', msg='multiline string equal failed')
        # 强制测试失败
        self.fail('this test will be failed')

    # 正则表达式匹配
    def test_regex(self):
        # assertRegex 使用第二个参数 search 第一个参数，匹配则通过
        self.assertRegex('hello world', 'hello')
        # assertNotRegex 使用第二个参数 search 第一个参数，不匹配则通过
        self.assertNotRegex('hello world', 'world!')

    def test_mybe_skipped(self):
        if not '':
            self.skipTest('unittest.skipTest() skipped')


class TestContainerMethods(unittest.TestCase):
    '''所有测试类继承自 unittest.TestCase'''

    # 比较两个序列是否包含相同的元素，忽略元素顺序和容器类型
    def test_count_same_elements_in_two_sequences(self):
        self.assertCountEqual([1, 2, 3, 4], (3, 2, 1, 5))

    # 比较两个序列是否相等, 包括元素顺序, 忽略容器类型
    def test_sequence_equal(self):
        self.assertSequenceEqual([1, 2, 3], (1, 2, 3))

    # 比较两个序列是否相等, 包括元素顺序和容器类型
    def test_sequence_with_seqtype_equal(self):
        self.assertSequenceEqual([1, 2, 3], (1, 2, 3), seq_type=list)

    # 比较两个列表是否相等, 包括元素顺序和容器类型
    def test_list_equal(self):
        self.assertListEqual([1, 2, 3], [3, 2, 1])

    # 比较两个元组是否相等, 包括元素顺序和容器类型
    def test_tuple_equal(self):
        self.assertTupleEqual((1, 2, 3), (3, 2, 1))

    # 比较两个集合(不可变集合)是否相等
    def test_set_equal(self):
        self.assertSetEqual({1, 2, 3}, {3, 2, 1})
        self.assertSetEqual(frozenset({3, 2, 1}), {1, 2, 3})

    # 比较两个字典是否相等
    def test_dict_equal(self):
        self.assertDictEqual({'a': 1, 'b': 2}, {'b': 2, 'a': 1})


# 自定义测试结果类 继承 unittest.TestResult
class CustomTestresult(unittest.TestResult):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.passed_tests = []
        self.failed_tests = []
        self.skipped_tests = []

    def addSuccess(self, test):
        '''测试通过时调用'''
        super().addSuccess(test)
        self.passed_tests.append(test)
        print(f'success: {test}')

    def addFailure(self, test, err):
        '''测试失败时调用'''
        super().addFailure(test, err)
        self.failed_tests.append(test)
        print(f'failure: {test}')

    def addSkip(self, test, reason):
        '''测试跳过时调用'''
        super().addSkip(test, reason)
        self.skipped_tests.append(test)
        print(f'skip: {test}')

    def get_report(self):
        '''返回测试报告'''
        report = []
        report.append(f'passed tests: {self.passed_tests}')
        report.append(f'failed tests: {self.failed_tests}')
        report.append(f'skipped tests: {self.skipped_tests}')
        return '\n'.join(report)


# 自定义测试运行器类 继承 unittest.TextTestRunner
class CustomTestRunner(unittest.TextTestRunner):
    def _makeResult(self):
        return CustomTestresult()


if __name__ == '__main__':
    # 方式1: 使用默认测试运行器
    # unittest.main()

    # 方式2: 自定义测试运行器
    runner = CustomTestRunner()
    suite = unittest.TestLoader().loadTestsFromTestCase(TestStringMethods)
    result = runner.run(suite)
    print(f'testResult: {result.get_report()}')
